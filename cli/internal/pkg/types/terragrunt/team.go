package terragrunt

import (
	"errors"
	"fmt"
	"gh_foundations/internal/pkg/types/terraform_state"

	"github.com/tidwall/gjson"
)

type TeamImportIdResolver struct {
	StateExplorer terraform_state.IStateExplorer
}

func (t *TeamImportIdResolver) ResolveImportId(resourceAddress string) (string, error) {
	name, err := t.StateExplorer.GetResourceChangeAfterAttribute(resourceAddress, "name")
	if err != nil {
		return "", err
	} else if !name.Exists() {
		return "", fmt.Errorf("unable to resolve import id: missing %q attribute", "name")
	}

	return name.String(), nil
}

type TeamMemberImportIdResolver struct {
	StateExplorer terraform_state.IStateExplorer
}

func (t *TeamMemberImportIdResolver) ResolveImportId(resourceAddress string) (string, error) {
	var allErrors error
	teamId, err := t.StateExplorer.GetResourceChangeAfterAttribute(resourceAddress, "team_id")
	if err != nil {
		allErrors = errors.Join(allErrors, err)
	} else if !teamId.Exists() {
		allErrors = errors.Join(allErrors, fmt.Errorf("unable to resolve import id: missing %q attribute", "team_id"))
	}

	username, err := t.StateExplorer.GetResourceChangeAfterAttribute(resourceAddress, "username")
	if err != nil {
		allErrors = errors.Join(allErrors, err)
	} else if !username.Exists() {
		allErrors = errors.Join(allErrors, fmt.Errorf("unable to resolve import id: missing %q attribute", "username"))
	}

	if teamId == nil {
		teamId = &gjson.Result{Type: gjson.String, Str: ""}
	}

	if username == nil {
		username = &gjson.Result{Type: gjson.String, Str: ""}
	}

	return fmt.Sprintf("%s:%s", teamId.String(), username.String()), allErrors
}
