package import_cmd

import (
	"fmt"
	"gh_foundations/internal/pkg/functions"
	types "gh_foundations/internal/pkg/types/terragrunt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tidwall/gjson"
)

type planAndArchiveMsg struct {
	archive           types.IPlanFile
	resourceAddresses []string
}

type terragruntImportMsg int

type resolveResourceIdMsg string

type errMsg struct{ err error }

func generatePlanFile(modulePath string) tea.Cmd {
	return func() tea.Msg {
		moduleDir := functions.GetTerragruntModuleDir(modulePath)
		planName := "import_plan"
		outputFilePath := moduleDir + string(os.PathSeparator) + planName + ".json"
		planArchive, err := types.NewTerragruntPlanFile(planName, modulePath, moduleDir, outputFilePath)
		if err != nil {
			return errMsg{err}
		}

		err = planArchive.RunPlan(nil)
		if err != nil {
			return errMsg{err}
		}

		stateExplorer, err := planArchive.GetStateExplorer()
		if err != nil {
			return errMsg{err}
		}

		addresses, err := stateExplorer.GetChangedResourceAddresses(func(json gjson.Result) bool {
			gjsonActions := json.Get("change.actions")
			if !gjsonActions.Exists() || !gjsonActions.IsArray() {
				return false
			}
			actions := gjsonActions.Array()
			return len(actions) == 1 && actions[0].Type == gjson.String && actions[0].String() == "create"
		})
		if err != nil {
			return errMsg{err}
		}

		return planAndArchiveMsg{archive: planArchive, resourceAddresses: addresses}
	}
}

func resolveResourceId(address string, archive types.IPlanFile) tea.Cmd {

	return func() tea.Msg {
		explorer, err := archive.GetStateExplorer()
		if err != nil {
			return errMsg{err}
		}
		idResolver := functions.CreateImportIdResolver(address, explorer)
		if idResolver != nil {
			id, err := idResolver.ResolveImportId(address)
			if err != nil {
				err = archive.RunPlan(&address)
				if err != nil {
					return errMsg{err}
				}

				err = explorer.SetPlanFile(archive.GetPlanFilePath())
				if err != nil {
					return errMsg{err}
				}

				idResolver = functions.CreateImportIdResolver(address, explorer)
				id, _ = idResolver.ResolveImportId(address)
			}
			return resolveResourceIdMsg(id)
		}
		return errMsg{fmt.Errorf("no import ID resolver found for resource %q", address)}
	}
}

func runTerragruntImport(modulePath string, address string, id string) tea.Cmd {
	return func() tea.Msg {
		errBytes, err := functions.RunImportCommand(modulePath, address, id)
		if err != nil {
			return errMsg{fmt.Errorf("error running import command: %s", errBytes.String())}
		}
		return terragruntImportMsg(0)
	}
}
