package functions

import (
	"bytes"
	"gh_foundations/internal/pkg/types/terraform_state"
	types "gh_foundations/internal/pkg/types/terragrunt"
	"os/exec"
	"path/filepath"
)

func GetTerragruntModuleDir(modulePath string) string {
	return filepath.Dir(modulePath)
}

func RunImportCommand(modulePath string, address string, id string) (bytes.Buffer, error) {
	moduleDir := GetTerragruntModuleDir(modulePath)
	errorBytes := bytes.Buffer{}
	importCmd := exec.Command("terragrunt", "import", address, id)
	importCmd.Stderr = &errorBytes
	importCmd.Stdout = nil
	importCmd.Dir = moduleDir
	return errorBytes, importCmd.Run()
}

func CreateImportIdResolver(resourceAddress string, stateExplorer terraform_state.IStateExplorer) types.ImportIdResolver {
	resourceType, err := stateExplorer.GetResourceChangeResourceType(resourceAddress)
	if err != nil {
		return nil
	}
	switch resourceType {
	case "github_team_membership":
		return &types.TeamMemberImportIdResolver{StateExplorer: stateExplorer}
	case "github_team":
		return &types.TeamImportIdResolver{StateExplorer: stateExplorer}
	case "github_repository":
		return &types.RepositoryImportIdResolver{StateExplorer: stateExplorer}
	case "github_branch_default":
		return &types.RepositoryBranchDefaultImportIdResolver{StateExplorer: stateExplorer}
	case "github_repository_collaborators":
		return &types.RepositoryCollaboratorsImportIdResolver{StateExplorer: stateExplorer}
	case "github_actions_secret", "github_codespaces_secret", "dependabot_secret":
		return &types.RepositorySecretsImportIdResolver{StateExplorer: stateExplorer}
	case "github_repository_dependabot_security_updates":
		return &types.RepositoryDependabotSecurityUpdatesImportIdResolver{StateExplorer: stateExplorer}
	case "github_repository_environment":
		return &types.RepositoryEnvironmentImportIdResolver{StateExplorer: stateExplorer}
	case "github_repository_ruleset":
		return &types.RepositoryRulesetImportIdResolver{StateExplorer: stateExplorer}
	default:
		return nil
	}
}
