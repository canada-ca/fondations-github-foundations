package functions

import (
	"fmt"
	githubfoundations "gh_foundations/internal/pkg/types/github_foundations"
	"regexp"

	"github.com/tidwall/gjson"
)

func MapTerraformerRepositoryToGithubFoundationRepository(rAttributes gjson.Result) *githubfoundations.RepositoryInput {
	var templateRepository *githubfoundations.TemplateRepositoryInputs
	if GjsonGetDefault(rAttributes, "template\\.\\#", 0, func(r gjson.Result) int { return int(r.Int()) }) > 0 {
		templateRepository = &githubfoundations.TemplateRepositoryInputs{
			Owner:              rAttributes.Get("template\\.\\0\\.owner").String(),
			Repository:         rAttributes.Get("template\\.\\0\\.repository").String(),
			IncludeAllBranches: GjsonGetDefault(rAttributes, "template\\.\\0\\.include_all_branches", false, func(r gjson.Result) bool { return r.Bool() }),
		}
	}

	return &githubfoundations.RepositoryInput{
		Name:                              rAttributes.Get("name").String(),
		Description:                       rAttributes.Get("description").String(),
		DefaultBranch:                     rAttributes.Get("default_branch").String(),
		RepositoryTeamPermissionsOverride: make(map[string]string),
		ProtectedBranches:                 make([]string, 0),
		AdvanceSecurity: GjsonGetDefault(rAttributes, "security_and_analysis\\.\\0\\.advanced_security\\.\\0\\.status", false, func(r gjson.Result) bool {
			return r.String() == "enabled"
		}),
		HasVulnerabilityAlerts:    rAttributes.Get("vulnerability_alerts").Bool(),
		Topics:                    GjsonGetList[string](rAttributes, "topics", func(r gjson.Result) string { return r.String() }),
		Homepage:                  rAttributes.Get("homepage").String(),
		DeleteHeadBranchOnMerge:   GjsonGetDefault(rAttributes, "delete_branch_on_merge", true, func(r gjson.Result) bool { return r.Bool() }),
		RequiresWebCommitSignOff:  GjsonGetDefault(rAttributes, "web_commit_signoff_required", true, func(r gjson.Result) bool { return r.Bool() }),
		DependabotSecurityUpdates: false,
		LicenseTemplate:           GjsonGetDefault(rAttributes, "license_template", "", func(r gjson.Result) string { return r.String() }),
		TemplateRepository:        templateRepository,
		AllowAutoMerge:            GjsonGetDefault(rAttributes, "allow_auto_merge", false, func(r gjson.Result) bool { return r.Bool() }),
	}
}

func GjsonGetDefault[T any](obj gjson.Result, key string, defaultValue T, conversion func(r gjson.Result) T) T {
	result := obj.Get(key)
	if result.Exists() {
		return conversion(result)
	}
	return defaultValue
}

func GjsonGetList[T any](obj gjson.Result, key string, conversion func(r gjson.Result) T) []T {
	count := obj.Get(fmt.Sprintf("%s\\.#", key)).Int()
	values := make([]T, count)
	for index := range values {
		values[index] = conversion(obj.Get(fmt.Sprintf("%s\\.%d", key, index)))
	}
	return values
}

func IdentifyFoundationsResourceType(resource_id string) githubfoundations.ResourceType {
	resourceRegexp := regexp.MustCompile("^[A-Za-z_]+")
	match := resourceRegexp.FindStringSubmatch(resource_id)
	if len(match) == 0 {
		return githubfoundations.None
	}

	switch match[0] {
	case "github_repository":
		return githubfoundations.Repository
	case "github_repository_collaborator":
		return githubfoundations.RepositoryCollaborator
	default:
		return githubfoundations.None
	}
}
