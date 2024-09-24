package terragrunt

import (
	"fmt"
	"gh_foundations/internal/pkg/types/terraform_state"

	"github.com/tidwall/gjson"
)

type RepositoryImportIdResolver struct {
	StateExplorer terraform_state.IStateExplorer
}

func (r *RepositoryImportIdResolver) ResolveImportId(resourceAddress string) (string, error) {
	name, err := r.StateExplorer.GetResourceChangeAfterAttribute(resourceAddress, "name")
	if err != nil {
		return "", err
	} else if !name.Exists() {
		return "", fmt.Errorf("unable to resolve import id: unexpected error occurred")
	}

	return name.String(), nil
}

type RepositoryBranchDefaultImportIdResolver struct {
	StateExplorer terraform_state.IStateExplorer
}

func (r *RepositoryBranchDefaultImportIdResolver) ResolveImportId(resourceAddress string) (string, error) {
	repository, err := r.StateExplorer.GetResourceChangeAfterAttribute(resourceAddress, "repository")
	if err != nil {
		return "", err
	} else if !repository.Exists() {
		return "", fmt.Errorf("unable to resolve import id: unexpected error occurred")
	}

	return repository.String(), nil
}

type RepositoryCollaboratorsImportIdResolver struct {
	StateExplorer terraform_state.IStateExplorer
}

func (r *RepositoryCollaboratorsImportIdResolver) ResolveImportId(resourceAddress string) (string, error) {
	repository, err := r.StateExplorer.GetResourceChangeAfterAttribute(resourceAddress, "repository")
	if err != nil {
		return "", err
	} else if !repository.Exists() {
		return "", fmt.Errorf("unable to resolve import id: unexpected error occurred")
	}

	return repository.String(), nil
}

type RepositorySecretsImportIdResolver struct {
	StateExplorer terraform_state.IStateExplorer
}

func (t *RepositorySecretsImportIdResolver) ResolveImportId(resourceAddress string) (string, error) {
	repository, err := t.StateExplorer.GetResourceChangeAfterAttribute(resourceAddress, "repository")
	if err != nil {
		return "", err
	} else if !repository.Exists() {
		return "", fmt.Errorf("unable to resolve import id: unexpected error occurred")
	} else if repository.Type != gjson.String {
		return "", fmt.Errorf("unable to resolve import id: repository attribute is not a string")
	}

	secretName, err := t.StateExplorer.GetResourceChangeAfterAttribute(resourceAddress, "secret_name")
	if err != nil {
		return "", err
	} else if !secretName.Exists() {
		return "", fmt.Errorf("unable to resolve import id: unexpected error occurred")
	} else if secretName.Type != gjson.String {
		return "", fmt.Errorf("unable to resolve import id: secret_name attribute is not a string")
	}

	return fmt.Sprintf("%s/%s", repository.String(), secretName.String()), nil
}

type RepositoryDependabotSecurityUpdatesImportIdResolver struct {
	StateExplorer terraform_state.IStateExplorer
}

func (t *RepositoryDependabotSecurityUpdatesImportIdResolver) ResolveImportId(resourceAddress string) (string, error) {
	repository, err := t.StateExplorer.GetResourceChangeAfterAttribute(resourceAddress, "repository")
	if err != nil {
		return "", err
	} else if !repository.Exists() {
		return "", fmt.Errorf("unable to resolve import id: unexpected error occurred")
	} else if repository.Type != gjson.String {
		return "", fmt.Errorf("unable to resolve import id: repository attribute is not a string")
	}

	return repository.String(), nil
}

type RepositoryEnvironmentImportIdResolver struct {
	StateExplorer terraform_state.IStateExplorer
}

func (t *RepositoryEnvironmentImportIdResolver) ResolveImportId(resourceAddress string) (string, error) {
	repository, err := t.StateExplorer.GetResourceChangeAfterAttribute(resourceAddress, "repository")
	if err != nil {
		return "", err
	} else if !repository.Exists() {
		return "", fmt.Errorf("unable to resolve import id: unexpected error occurred")
	} else if repository.Type != gjson.String {
		return "", fmt.Errorf("unable to resolve import id: repository attribute is not a string")
	}

	environment, err := t.StateExplorer.GetResourceChangeAfterAttribute(resourceAddress, "environment")
	if err != nil {
		return "", err
	} else if !environment.Exists() {
		return "", fmt.Errorf("unable to resolve import id: unexpected error occurred")
	} else if environment.Type != gjson.String {
		return "", fmt.Errorf("unable to resolve import id: environment attribute is not a string")
	}

	return fmt.Sprintf("%s/%s", repository.String(), environment.String()), nil
}

type RepositoryRulesetImportIdResolver struct {
	StateExplorer terraform_state.IStateExplorer
}

func (t *RepositoryRulesetImportIdResolver) ResolveImportId(resourceAddress string) (string, error) {
	repository, err := t.StateExplorer.GetResourceChangeAfterAttribute(resourceAddress, "repository")
	if err != nil {
		return "", err
	} else if !repository.Exists() {
		return "", fmt.Errorf("unable to resolve import id: unexpected error occurred")
	} else if repository.Type != gjson.String {
		return "", fmt.Errorf("unable to resolve import id: repository attribute is not a string")
	}

	// The full import id includes the ruleset id. We won't be able to get this info from a terraform plan. And will need to be typed out unless we add a github client to this tool.
	return fmt.Sprintf("%s:", repository.String()), nil
}
