package terraform_state

import (
	"errors"

	"github.com/tidwall/gjson"
)

var ErrUnknownAttribute = errors.New("attribute won't be known until after apply")
var ErrChangeAttributeNotFound = errors.New("attribute not found in change")

type IStateExplorer interface {
	GetChangedResourceAddresses(filterFn func(json gjson.Result) bool) ([]string, error)
	GetResourceChangeAfterAttribute(address string, attribute string) (*gjson.Result, error)
	GetResourceChangeResourceType(address string) (string, error)
	SetPlan(plan []byte)
	SetPlanFile(planFilePath string) error
}
