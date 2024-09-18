package githubfoundations

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

//Module Root Inputs

type RepositorySetInput struct {
	PrivateRepositories              []*RepositoryInput
	PublicRepositories               []*RepositoryInput
	DefaultRepositoryTeamPermissions map[string]string
}

func (r *RepositorySetInput) WriteHCL(file *hclwrite.File) {
	rootBody := file.Body()
	rootBodyMap := make(map[string]cty.Value)

	privateRepositories := make(map[string]cty.Value)
	for _, repository := range r.PrivateRepositories {
		privateRepositories[repository.Name] = repository.GetCtyValue()
	}

	publicRepositories := make(map[string]cty.Value)
	for _, repository := range r.PublicRepositories {
		publicRepositories[repository.Name] = repository.GetCtyValue()
	}

	rootBodyMap["private_repositories"] = cty.ObjectVal(privateRepositories)
	rootBodyMap["public_repositories"] = cty.ObjectVal(publicRepositories)
	rootBody.SetAttributeValue("inputs", cty.ObjectVal(rootBodyMap))
}

// Repository Inputs

type RepositoryInput struct {
	Name string
	// Required
	Description                       string
	DefaultBranch                     string
	RepositoryTeamPermissionsOverride map[string]string
	ProtectedBranches                 []string
	AdvanceSecurity                   bool
	HasVulnerabilityAlerts            bool
	Topics                            []string
	Homepage                          string
	DeleteHeadBranchOnMerge           bool
	RequiresWebCommitSignOff          bool
	DependabotSecurityUpdates         bool
	AllowAutoMerge                    bool
	// Optional
	OrganizationActionSecrets     []string
	OrganizationCodespaceSecrets  []string
	OrganizationDependabotSecrets []string
	ActionSecrets                 map[string]string
	CodespaceSecrets              map[string]string
	DependabotSecrets             map[string]string
	Environments                  map[string]EnvironmentInputs
	TemplateRepository            *TemplateRepositoryInputs
	LicenseTemplate               string
	UserPermissions               map[string]string
}

func (r *RepositoryInput) GetCtyValue() cty.Value {
	mapVal := make(map[string]cty.Value)

	// Required fields
	mapVal["description"] = cty.StringVal(r.Description)
	mapVal["default_branch"] = cty.StringVal(r.DefaultBranch)
	mapVal["advance_security"] = cty.BoolVal(r.AdvanceSecurity)
	mapVal["has_vulnerability_alerts"] = cty.BoolVal(r.HasVulnerabilityAlerts)
	mapVal["topics"] = toCtyValueSlice(r.Topics)
	mapVal["homepage"] = cty.StringVal(r.Homepage)
	mapVal["delete_head_on_merge"] = cty.BoolVal(r.DeleteHeadBranchOnMerge)
	mapVal["requires_web_commit_signing"] = cty.BoolVal(r.RequiresWebCommitSignOff)
	mapVal["dependabot_security_updates"] = cty.BoolVal(r.DependabotSecurityUpdates)
	mapVal["protected_branches"] = toCtyValueSlice(r.ProtectedBranches)
	mapVal["allow_auto_merge"] = cty.BoolVal(r.AllowAutoMerge)

	// Optional fields
	if len(r.OrganizationActionSecrets) > 0 {
		mapVal["organization_action_secrets"] = toCtyValueSlice(r.OrganizationActionSecrets)
	}

	if len(r.OrganizationCodespaceSecrets) > 0 {
		mapVal["organization_codespace_secrets"] = toCtyValueSlice(r.OrganizationCodespaceSecrets)
	}

	if len(r.OrganizationDependabotSecrets) > 0 {
		mapVal["organization_dependabot_secrets"] = toCtyValueSlice(r.OrganizationDependabotSecrets)
	}

	if len(r.ActionSecrets) > 0 {
		actionSecretsMap := make(map[string]cty.Value)
		for key, val := range r.ActionSecrets {
			actionSecretsMap[key] = cty.StringVal(val)
		}
		mapVal["action_secrets"] = cty.MapVal(actionSecretsMap)
	}
	if len(r.CodespaceSecrets) > 0 {
		codespaceSecretsMap := make(map[string]cty.Value)
		for key, val := range r.CodespaceSecrets {
			codespaceSecretsMap[key] = cty.StringVal(val)
		}
		mapVal["codespace_secrets"] = cty.MapVal(codespaceSecretsMap)
	}
	if len(r.DependabotSecrets) > 0 {
		dependabotSecretsMap := make(map[string]cty.Value)
		for key, val := range r.DependabotSecrets {
			dependabotSecretsMap[key] = cty.StringVal(val)
		}
		mapVal["dependabot_secrets"] = cty.MapVal(dependabotSecretsMap)
	}
	if len(r.Environments) > 0 {
		environmentsMap := make(map[string]cty.Value)
		for key, val := range r.Environments {
			environmentMap := make(map[string]cty.Value)
			actionSecretsMap := make(map[string]cty.Value)
			for key, val := range val.ActionSecrets {
				actionSecretsMap[key] = cty.StringVal(val)
			}
			environmentMap["action_secrets"] = cty.MapVal(actionSecretsMap)
			environmentsMap[key] = cty.ObjectVal(environmentMap)
		}
		mapVal["environments"] = cty.MapVal(environmentsMap)
	}

	if r.TemplateRepository != nil {
		templateRepoMap := make(map[string]cty.Value)
		templateRepoMap["owner"] = cty.StringVal(r.TemplateRepository.Owner)
		templateRepoMap["repository"] = cty.StringVal(r.TemplateRepository.Repository)
		templateRepoMap["include_all_branches"] = cty.BoolVal(r.TemplateRepository.IncludeAllBranches)
		mapVal["template_repository"] = cty.ObjectVal(templateRepoMap)
	}

	if r.LicenseTemplate != "" {
		mapVal["license_template"] = cty.StringVal(r.LicenseTemplate)
	}

	if len(r.UserPermissions) > 0 {
		userPermissionsMap := make(map[string]cty.Value)
		for key, val := range r.UserPermissions {
			userPermissionsMap[key] = cty.StringVal(val)
		}
		mapVal["user_permissions"] = cty.MapVal(userPermissionsMap)
	}
	return cty.ObjectVal(mapVal)
}

type EnvironmentInputs struct {
	ActionSecrets map[string]string
}

type TemplateRepositoryInputs struct {
	Owner              string `yaml:"Owner"`
	Repository         string `yaml:"Repository"`
	IncludeAllBranches bool   `yaml:"IncludeAllBranches"`
}
