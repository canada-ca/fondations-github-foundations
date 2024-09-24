package v1_2

import (
	"fmt"
	"gh_foundations/internal/pkg/types/terraform_state"
	"os"

	"github.com/tidwall/gjson"
)

type StateExplorer struct {
	parsedPlan gjson.Result
}

func (e *StateExplorer) GetChangedResourceAddresses(filterFn func(json gjson.Result) bool) ([]string, error) {
	query := "resource_changes"
	addresses := make([]string, 0)

	queryResult := e.parsedPlan.Get(query)
	if !queryResult.Exists() {
		return addresses, fmt.Errorf("no resource changes found in plan")
	}

	if !queryResult.IsArray() {
		return addresses, fmt.Errorf("resource changes query result is not an array")
	}

	for _, change := range queryResult.Array() {
		if filterFn(change) {
			address := change.Get("address")
			addresses = append(addresses, address.String())
		}
	}
	return addresses, nil
}
func (e *StateExplorer) GetResourceChangeAfterAttribute(address string, attribute string) (*gjson.Result, error) {
	query := fmt.Sprintf("resource_changes.#(address==%q).change.after.%s", address, gjson.Escape(attribute))
	result := e.parsedPlan.Get(query)
	if !result.Exists() {
		unknownQuery := fmt.Sprintf("resource_changes.#(address==%q).change.after_unknown.%s", address, gjson.Escape(attribute))
		result = e.parsedPlan.Get(unknownQuery)
		if result.Exists() && result.Bool() {
			return nil, terraform_state.ErrUnknownAttribute
		}
		return nil, terraform_state.ErrChangeAttributeNotFound
	}

	return &result, nil
}

func (e *StateExplorer) GetResourceChangeResourceType(address string) (string, error) {
	query := fmt.Sprintf("resource_changes.#(address==%q).type", address)
	result := e.parsedPlan.Get(query)
	if !result.Exists() {
		return "", fmt.Errorf("resource type not found for address %q", address)
	}

	return result.String(), nil
}

func (e *StateExplorer) SetPlan(plan []byte) {
	e.parsedPlan = gjson.ParseBytes(plan)
}

func (e *StateExplorer) SetPlanFile(planFilePath string) error {
	bytes, err := os.ReadFile(planFilePath)
	if err != nil {
		return err
	}
	e.parsedPlan = gjson.ParseBytes(bytes)
	return nil
}
