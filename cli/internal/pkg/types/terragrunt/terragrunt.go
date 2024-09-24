package terragrunt

import (
	"bytes"
	"fmt"
	"gh_foundations/internal/pkg/types"
	"gh_foundations/internal/pkg/types/status"
	"gh_foundations/internal/pkg/types/terraform_state"
	v1_2 "gh_foundations/internal/pkg/types/terraform_state/v1.2"
	"io"
	"log"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

var fs = afero.NewOsFs()

// command creation function for mocking
var newCommandExecutor = func(name string, args ...string) types.ICommandExecutor {
	return &types.CommandExecutor{
		Cmd: exec.Command(name, args...),
	}
}

type IPlanFile interface {
	Cleanup() error
	RunPlan(target *string) error
	GetStateExplorer() (terraform_state.IStateExplorer, error)
	GetPlanFilePath() string
}

type PlanFile struct {
	Name           string
	ModulePath     string
	ModuleDir      string
	OutputFilePath string
}

type HCLFile struct {
	Path string
}

func getRepository(repo map[string]interface{}) (status.Repository, error) {
	var repository status.Repository

	err := mapstructure.Decode(repo, &repository)
	if err != nil {
		log.Fatalf("Error in getInputsFromFile mapstructure.Decode: %s", err)
		return repository, err
	}

	return repository, nil

}


// Given a repository map, returned by Viper, return a map of status.Repository
func getRepositoryMap(repoList []map[string]interface{}) (map[string]status.Repository, error) {
	repos := make(map[string]status.Repository)

	for name, r := range repoList[0] {
		details := r.([]map[string]interface{})
		d := details[0]
		repo, err := getRepository(d)
		if err != nil {
			log.Fatalf("Error in getRepositoryMap: %s", err)
			return repos, err
		}
		repos[name] = repo
	}
	return repos, nil
}

// Return the locals block from the HCL file as a slice of string slices
func getLocalsBlock(contents string) [][]string {
	// The locals are in the form of locals = { key = value }

	// Use regex to find the locals block
	lre := regexp.MustCompile(`locals\s*{\n*((.*[^}])\n)+}`)
	locals := lre.FindString(contents)

	if locals == "" {
		fmt.Printf("locals not found")
		return make([][]string, 0)
	}

	// Use regex to find the key-value pairs in the locals block
	kvre := regexp.MustCompile(`(.*[^=])=(.*(?:(:?\n.*[^\]])*])*)`)
	matches := kvre.FindAllStringSubmatch(locals, -1)

	// Clean up the matches a little
	for i, match := range matches {
		matches[i][1] = strings.Trim(match[1], " \"")
		matches[i][2] = strings.Trim(match[2], " \"")
	}

	return matches
}


// The locals are in the form of locals = { key = value }
// Then, they are referred to as local.key in the configuration
// This function replaces the locals with their values
func replaceLocals(contents string) string {
	matches := getLocalsBlock(contents)

	// Replace the locals with their values
	for _, match := range matches {
		key := strings.Trim(match[1], " ")
		value := strings.Trim(match[2], " ")

		contents = strings.ReplaceAll(contents, key, value)

	}

    return contents
}

// Given an HCL file, return the inputs
func (h *HCLFile) GetInputsFromFile() (status.Inputs, error) {

	var inputs status.Inputs

	viper.SetConfigType("hcl")
	viper.SetConfigFile(h.Path)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found;
			log.Fatalf(`GetInputsFromFile: config file not found: %s`, h.Path)
			return inputs, err
		} else if _, ok := err.(viper.ConfigParseError); ok {
			// the viper library can't parse the "locals". Here's a workaround
			// read in the file contents and replace the locals with their values
			// then write the file to a temporary location and read it in
			// again

			// read in the file contents
			contents, err := afero.ReadFile(fs, h.Path)
			if err != nil {
				log.Fatalf(`GetInputsFromFile: unable to read config file: %s`, h.Path)
				return inputs, err
			}
			// replace the locals with their values
			contents = []byte(replaceLocals(string(contents)))
			// write the file to a temporary location and read it in again
			tempPath := "/tmp/" + path.Base(h.Path)
			err = afero.WriteFile(fs, tempPath, contents, 0644)
			if err != nil {
				log.Fatalf(`GetInputsFromFile: unable to write config file: %s`, h.Path)
				return inputs, err
			}
			viper.SetConfigFile(tempPath)
			viper.ReadInConfig()
		} else {
			// Config file was found but another error was produced
			log.Fatalf(`GetInputsFromFile: config file found but another error was produced: %s`, h.Path)
			return inputs, err
		}
	}

	raw := viper.Get("inputs").([]map[string]interface{})
	for key, input := range raw[0] {
		switch key {
			case "private_repositories":
				repoList := input.([]map[string]interface{})
				repos, err := getRepositoryMap(repoList)
				if err != nil {
					log.Fatalf("Error in getRepositoryMap: %s", err)
					return inputs, err
				}
				inputs.PrivateRepositories = repos
			case "public_repositories":
				repoList := input.([]map[string]interface{})
				repos, err := getRepositoryMap(repoList)
				if err != nil {
					log.Fatalf("Error in getRepositoryMap: %s", err)
					return inputs, err
				}
				inputs.PublicRepositories = repos
			case "default_repository_team_permissions":
				permissions := make(map[string]string)
				inputArr := input.([]map[string]interface{})
				drtpsArr := inputArr[0]
				for permission, value := range drtpsArr {
					permissions[permission] = value.(string)
				}
				inputs.DefaultRepositoryTeamPermissions = permissions
			default:
				log.Fatalf("Unknown input: %s", key)
				return inputs, nil
		}
	}

	return inputs, nil
}


// Given the string content of an HCL file, return a map of locals
func (h *HCLFile) GetLocalsMap() map[string]string {

	// If the path is not set, return an empty map
	if h.Path == "" {
		return make(map[string]string)
	}

	// If the path is set, read the file and return the locals
	content, err := afero.ReadFile(fs, h.Path)
	if err != nil {
		log.Fatalf(`GetLocalsMap: unable to read config file: %s`, h.Path)
		return make(map[string]string)
	}

	macthes := getLocalsBlock(string(content))
	locals := make(map[string]string)
	for _, match := range macthes {
		locals[match[1]] = match[2]
	}
	return locals
}

func NewTerragruntPlanFile(name string, modulePath string, moduleDir string, outputFilePath string) (*PlanFile, error) {
	// If there is a file conflict with the output file, create a new file with a "copy_" prefix
	if _, err := fs.Stat(outputFilePath); err == nil {
		dir := path.Dir(outputFilePath)
		filename := path.Base(outputFilePath)
		outputFilePath = path.Join(dir, "copy_"+filename)
	}

	return &PlanFile{
		Name:           name,
		ModuleDir:      moduleDir,
		ModulePath:     modulePath,
		OutputFilePath: outputFilePath,
	}, nil
}

func (t *PlanFile) Cleanup() error {
	return fs.Remove(t.OutputFilePath)
}

func (t *PlanFile) GetPlanFilePath() string {
	return t.OutputFilePath
}

func (t *PlanFile) RunPlan(target *string) error {
	if _, errBytes, err := runPlan(t.ModuleDir, &t.Name, target); err != nil {
		return fmt.Errorf("error running plan: %s", errBytes.String())
	}

	planFile, err := fs.Create(t.OutputFilePath)
	if err != nil {
		return err
	}
	defer planFile.Close()

	if errBytes, err := outputPlan(t.Name, planFile, t.ModuleDir); err != nil {
		return fmt.Errorf("error outputting plan: %s", errBytes.String())
	}

	return nil
}

func (t *PlanFile) GetStateExplorer() (terraform_state.IStateExplorer, error) {
	planBytes, err := afero.ReadFile(fs, t.OutputFilePath)
	if err != nil {
		return nil, err
	}

	var explorer terraform_state.IStateExplorer
	versionQuery := "format_version"
	gjsonResult := gjson.GetBytes(planBytes, versionQuery)
	if !gjsonResult.Exists() {
		return nil, fmt.Errorf("unable to determine plan version")
	} else if gjsonResult.Type != gjson.String {
		return nil, fmt.Errorf("unexpected type for %q: %s", versionQuery, gjsonResult.Type)
	}
	version := gjsonResult.String()

	switch version {
	case "1.2":
		explorer = &v1_2.StateExplorer{}
	default:
		return nil, fmt.Errorf("unsupported version %q", version)
	}

	explorer.SetPlan(planBytes)
	return explorer, nil
}

type ImportIdResolver interface {
	ResolveImportId(resourceAddress string) (string, error)
}

func outputPlan(planName string, planFile io.Writer, dir string) (bytes.Buffer, error) {
	errBuffer := &bytes.Buffer{}
	cmdExecutor := newCommandExecutor("terragrunt", "show", "-json", planName)
	cmdExecutor.SetOutput(planFile)
	cmdExecutor.SetErrorOutput(errBuffer)
	cmdExecutor.SetDir(dir)
	if err := cmdExecutor.Run(); err != nil {
		return *errBuffer, err
	}
	return *errBuffer, nil
}

func runPlan(dir string, output *string, target *string) (bytes.Buffer, bytes.Buffer, error) {
	errBuffer := &bytes.Buffer{}
	logBuffer := &bytes.Buffer{}
	args := []string{"plan", "-lock=false"}
	if output != nil {
		args = append(args, fmt.Sprintf("-out=%s", *output))
	}
	if target != nil {
		args = append(args, fmt.Sprintf("-target=%s", *target))
	}

	cmdExecutor := newCommandExecutor("terragrunt", args...)
	cmdExecutor.SetErrorOutput(errBuffer)
	cmdExecutor.SetDir(dir)
	if err := cmdExecutor.Run(); err != nil {
		return *logBuffer, *errBuffer, err
	} else {
		logBuffer.WriteString(fmt.Sprintf("Command %q complete", cmdExecutor.String()))
	}
	return *logBuffer, *errBuffer, nil
}
