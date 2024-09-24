package githubfoundations

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// Module Root Inputs

type TeamSetInput struct {
	Teams []*TeamInput
}

func (r *TeamSetInput) WriteHCL(file *hclwrite.File) {
	rootBody := file.Body()
	rootBodyMap := make(map[string]cty.Value)

	teams := make(map[string]cty.Value)
	for _, team := range r.Teams {
		teams[team.Name] = team.GetCtyValue()
	}

	rootBodyMap["teams"] = cty.ObjectVal(teams)
	rootBody.SetAttributeValue("inputs", cty.ObjectVal(rootBodyMap))
}

type TeamInput struct {
	Name        string
	Description string
	Privacy     string
	Maintainers []string
	Members     []string
	ParentId    string
}

func (t *TeamInput) GetCtyValue() cty.Value {
	mapVal := make(map[string]cty.Value)
	mapVal["description"] = cty.StringVal(t.Description)
	mapVal["privacy"] = cty.StringVal(t.Privacy)

	if len(t.ParentId) > 0 {
		mapVal["parent_id"] = cty.StringVal(t.ParentId)
	}

	mapVal["maintainers"] = toCtyValueSlice(t.Maintainers)
	mapVal["members"] = toCtyValueSlice(t.Members)

	return cty.ObjectVal(mapVal)
}
