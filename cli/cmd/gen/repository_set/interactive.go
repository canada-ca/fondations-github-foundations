package repositoryset

import (
	"fmt"
	"gh_foundations/cmd/gen/common"
	githubfoundations "gh_foundations/internal/pkg/types/github_foundations"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	yaml "gopkg.in/yaml.v2"
)

var questions []common.IQuestion = []common.IQuestion{
	common.NewSelectQuestion(
		"Select the visibility of the repository",
		[]string{
			"public",
			"private",
		},
	),
	common.NewTextQuestion(
		"Enter the name of the repository",
		"",
	),
	common.NewTextQuestion(
		"Enter the description for the repository",
		"",
	),
	common.NewTextQuestion(
		"Enter the default branch for the repository",
		"main",
	),
	common.NewListQuestion(
		"Enter the name(s) of any protected branches",
	),
	common.NewKeyValueListQuestion(
		"Enter custom team permissions for the repository",
	),
	common.NewKeyValueListQuestion(
		"Enter custom user permissions for the repository",
	),
	common.NewSelectQuestion(
		"Enable Github Advance Security",
		[]bool{
			true,
			false,
		},
	),
	common.NewSelectQuestion(
		"Enable vulnerability alerts",
		[]bool{
			true,
			false,
		},
	),
	common.NewListQuestion(
		"Add Topics",
	),
	common.NewTextQuestion(
		"Enter the homepage for the repository",
		"",
	),
	common.NewSelectQuestion(
		"Delete head branches on merge",
		[]bool{
			true,
			false,
		},
	),
	common.NewSelectQuestion(
		"Require web commit signoff",
		[]bool{
			true,
			false,
		},
	),
	common.NewSelectQuestion(
		"Enable Dependabot security updates",
		[]bool{
			true,
			false,
		},
	),
	common.NewSelectQuestion(
		"Allow auto merge",
		[]bool{
			true,
			false,
		},
	),
	common.NewCompositeQuestion(
		"Fill out the following to create the repository using a template repository",
		[]common.CompositeQuestionEntry{
			{
				Key: "Owner",
				Question: common.NewTextQuestion(
					"Enter the owner of the template repository",
					"",
				),
			},
			{
				Key: "Repository",
				Question: common.NewTextQuestion(
					"Enter the name of the template repository",
					"",
				),
			}, {
				Key: "IncludeAllBranches",
				Question: common.NewSelectQuestion(
					"Include all branches from template repository",
					[]bool{
						true,
						false,
					},
				),
			},
		},
	),
	common.NewTextQuestion(
		"Enter the name of a license template",
		"",
	),
}

func submitFunc(answers []string, repositorySet *githubfoundations.RepositorySetInput) {
	repository := new(githubfoundations.RepositoryInput)
	repository.Name = answers[1]
	repository.Description = answers[2]
	repository.DefaultBranch = answers[3]

	protectedBranches := make([]string, 0)
	err := yaml.Unmarshal([]byte(answers[4]), &protectedBranches)
	if err != nil {
		fmt.Println("Error converting protected branches input to array of strings:", err)
		os.Exit(1)
	}
	repository.ProtectedBranches = protectedBranches

	teamPermissionOverrides := make(map[string]string, 0)
	err = yaml.Unmarshal([]byte(answers[5]), &teamPermissionOverrides)
	if err != nil {
		fmt.Println("Error converting team permissions input to map of strings:", err)
		os.Exit(1)
	}
	repository.RepositoryTeamPermissionsOverride = teamPermissionOverrides

	userPermissions := make(map[string]string, 0)
	err = yaml.Unmarshal([]byte(answers[6]), &userPermissions)
	if err != nil {
		fmt.Println("Error converting user permissions input to map of strings:", err)
		os.Exit(1)
	}
	repository.UserPermissions = userPermissions

	var enableAdvanceSecurity bool
	err = yaml.Unmarshal([]byte(answers[7]), &enableAdvanceSecurity)
	if err != nil {
		fmt.Println("Error converting advance security input to boolean:", err)
		os.Exit(1)
	}
	repository.AdvanceSecurity = enableAdvanceSecurity

	var enableVulnerabilityAlerts bool
	err = yaml.Unmarshal([]byte(answers[8]), &enableVulnerabilityAlerts)
	if err != nil {
		fmt.Println("Error converting vulnerability alerts input to boolean:", err)
		os.Exit(1)
	}
	repository.HasVulnerabilityAlerts = enableVulnerabilityAlerts

	topics := make([]string, 0)
	err = yaml.Unmarshal([]byte(answers[9]), &topics)
	if err != nil {
		fmt.Println("Error converting topics input to array of strings:", err)
		os.Exit(1)
	}
	repository.Topics = topics

	repository.Homepage = answers[10]

	var deleteHeadOnMerge bool
	err = yaml.Unmarshal([]byte(answers[11]), &deleteHeadOnMerge)
	if err != nil {
		fmt.Println("Error converting delete head branch on merge input to boolean:", err)
		os.Exit(1)
	}
	repository.DeleteHeadBranchOnMerge = deleteHeadOnMerge

	var requireWebCommitSignoff bool
	err = yaml.Unmarshal([]byte(answers[12]), &requireWebCommitSignoff)
	if err != nil {
		fmt.Println("Error converting require web commit signoff on merge input to boolean:", err)
		os.Exit(1)
	}
	repository.RequiresWebCommitSignOff = requireWebCommitSignoff

	var enableDependabotSecurityUpdates bool
	err = yaml.Unmarshal([]byte(answers[13]), &enableDependabotSecurityUpdates)
	if err != nil {
		fmt.Println("Error converting dependabot security updates input to boolean:", err)
		os.Exit(1)
	}
	repository.DependabotSecurityUpdates = enableDependabotSecurityUpdates

	var allowAutoMerge bool
	err = yaml.Unmarshal([]byte(answers[14]), &allowAutoMerge)
	if err != nil {
		fmt.Println("Error converting allow auto merge input to boolean:", err)
		os.Exit(1)
	}
	repository.AllowAutoMerge = allowAutoMerge

	var templateRepository githubfoundations.TemplateRepositoryInputs
	err = yaml.Unmarshal([]byte(answers[15]), &templateRepository)
	if err != nil {
		fmt.Printf("%+v", answers[15])
		fmt.Println("Error converting template repository input to object:", err)
		os.Exit(1)
	} else if len(templateRepository.Owner) > 0 && len(templateRepository.Repository) > 0 {
		repository.TemplateRepository = &templateRepository
	}

	repository.LicenseTemplate = answers[16]

	switch answers[0] {
	case "private":
		repositorySet.PrivateRepositories = append(repositorySet.PrivateRepositories, repository)
	case "public":
		repositorySet.PublicRepositories = append(repositorySet.PublicRepositories, repository)
	}
}

func runInteractive() (*githubfoundations.RepositorySetInput, error) {
	m := common.NewModel(questions, new(githubfoundations.RepositorySetInput), submitFunc)
	if _, err := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion()).Run(); err != nil {
		return nil, err
	}
	return m.Result, nil
}
