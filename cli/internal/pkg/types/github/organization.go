package github

import (
	"errors"
	"gh_foundations/internal/pkg/types"
	"slices"
	"time"

	"github.com/google/go-github/v61/github"
)

type Organization struct {
	*github.Organization
	customRepositoryRoles []github.CustomRepoRoles
}

func (o *Organization) Check(checkTypes []types.CheckType) types.CheckReport {
	report := types.CheckReport{
		EntityType: "github_organization",
		EntityId:   o.GetLogin(),
		Timestamp:  time.Now().Format(time.RFC3339), // Syslog compliant timestamp
		Results:    make(map[types.CheckType]types.CheckResult),
		Errors:     []types.CheckError{},
	}
	for _, t := range checkTypes {
		switch t {
		case types.GoCGuardrails:
			r, err := o.GoCGuardrailsCompliant()
			if err != nil {
				report.Errors = append(report.Errors, *err)
			}
			report.Results[t] = r
		default:
			report.Results[t] = types.NotApplicable
		}
	}
	return report
}

// Custom repository roles for an organization need to be accessed separately from settings
func (o *Organization) GoCGuardrailsCompliant() (types.CheckResult, *types.CheckError) {
	var allErrors error
	violations := make(map[string]string)
	checks := []func(org *Organization) (string, error){
		func(org *Organization) (string, error) {
			key := "dependabot_alerts_enabled_for_new_repositories"
			if !org.GetDependabotAlertsEnabledForNewRepos() {
				return key, errors.New("dependabot_alerts_enabled_for_new_repositories is not enabled. Expected it to be enabled")
			}
			return key, nil
		},
		func(org *Organization) (string, error) {
			key := "dependabot_security_updates_enabled_for_new_repositories"
			if !org.GetDependabotSecurityUpdatesEnabledForNewRepos() {
				return key, errors.New("dependabot_security_updates_enabled_for_new_repositories is not enabled. Expected it to be enabled")
			}
			return key, nil
		},
		func(org *Organization) (string, error) {
			key := "dependency_graph_enabled_for_new_repositories"
			if !org.GetDependencyGraphEnabledForNewRepos() {
				return key, errors.New("dependency_graph_enabled_for_new_repositories is not enabled. Expected it to be enabled")
			}
			return key, nil
		},
		func(org *Organization) (string, error) {
			key := "security_and_analysis_enabled_for_new_repositories"
			if !org.GetSecretScanningEnabledForNewRepos() {
				return key, errors.New("secret_scanning_enabled_for_new_repositories is not enabled. Expected it to be enabled")
			}
			return key, nil
		},
		func(org *Organization) (string, error) {
			key := "secret_scanning_push_protection_enabled_for_new_repositories"
			if !org.GetSecretScanningPushProtectionEnabledForNewRepos() {
				return key, errors.New("secret_scanning_push_protection_enabled_for_new_repositories is not enabled. Expected it to be enabled")
			}
			return key, nil
		},
		func(org *Organization) (string, error) {
			key := "delete_branch_on_merge"
			if org.GetMembersCanCreatePublicRepos() {
				return key, errors.New("members_can_create_public_repositories is enabled. Expected it to be disabled")
			}
			return key, nil
		},
		func(org *Organization) (string, error) {
			key := "members_can_create_private_repositories"
			if !org.GetMembersCanCreatePrivateRepos() {
				return key, errors.New("members_can_create_private_repositories is not enabled. Expected it to be enabled")
			}
			return key, nil
		},
		func(org *Organization) (string, error) {
			key := "members_can_create_internal_repositories"
			if !org.GetMembersCanCreateInternalRepos() {
				return key, errors.New("members_can_create_internal_repositories is not enabled. Expected it to be enabled")
			}
			return key, nil
		},
		func(org *Organization) (string, error) {
			key := "members_can_fork_private_repositories"
			if org.GetMembersCanForkPrivateRepos() {
				return key, errors.New("members_can_fork_private_repositories is enabled. Expected it to be disabled")
			}
			return key, nil
		},
		func(org *Organization) (string, error) {
			key := "security_engineer_role"
			for _, role := range org.customRepositoryRoles {
				base := role.GetBaseRole()
				_, deleteCodeScanningPerm := slices.BinarySearch(role.Permissions, "delete_alerts_code_scanning")
				_, writeCodeScanningPerm := slices.BinarySearch(role.Permissions, "write_code_scanning")
				if base == "maintain" && deleteCodeScanningPerm && writeCodeScanningPerm {
					return key, nil
				}
			}
			return key, errors.New("security engineer role undefined in the organization")
		},
		func(org *Organization) (string, error) {
			key := "contractor_role"
			for _, role := range org.customRepositoryRoles {
				base := role.GetBaseRole()
				_, manageWebhhooksPerm := slices.BinarySearch(role.Permissions, "manage_webhooks")
				if base == "write" && manageWebhhooksPerm {
					return key, nil
				}
			}
			return key, errors.New("contractor role undefined in the organization")
		},

		func(org *Organization) (string, error) {
			key := "community_manager_role"
			for _, role := range org.customRepositoryRoles {
				base := role.GetBaseRole()
				_, markAsDuplicatePerm := slices.BinarySearch(role.Permissions, "mark_as_duplicate")
				_, manageSettingsPagePerm := slices.BinarySearch(role.Permissions, "manage_settings_pages")
				_, manageSettingsWikiPerm := slices.BinarySearch(role.Permissions, "manage_settings_wiki")
				_, setSocialPreviewPrem := slices.BinarySearch(role.Permissions, "set_social_preview")
				_, editRepoMetdataPerm := slices.BinarySearch(role.Permissions, "edit_repo_metadata")
				_, editDiscussionCategoryPerm := slices.BinarySearch(role.Permissions, "edit_discussion_category")
				_, createDiscussionCategoryPerm := slices.BinarySearch(role.Permissions, "create_discussion_category")
				_, editCategoryOnDiscussionPerm := slices.BinarySearch(role.Permissions, "edit_category_on_discussion")
				_, toggleDiscussionAnswerPerm := slices.BinarySearch(role.Permissions, "toggle_discussion_answer")
				_, convertIssuesToDiscussionsPerm := slices.BinarySearch(role.Permissions, "convert_issues_to_discussions")
				_, closeDiscussionPerm := slices.BinarySearch(role.Permissions, "close_discussion")
				_, reopenDiscussionPerm := slices.BinarySearch(role.Permissions, "reopen_discussion")
				_, deleteDiscussionCommentPerm := slices.BinarySearch(role.Permissions, "delete_discussion_comment")
				if base == "read" && markAsDuplicatePerm && manageSettingsPagePerm && manageSettingsWikiPerm && setSocialPreviewPrem && editRepoMetdataPerm && editDiscussionCategoryPerm && createDiscussionCategoryPerm && editCategoryOnDiscussionPerm && toggleDiscussionAnswerPerm && convertIssuesToDiscussionsPerm && closeDiscussionPerm && reopenDiscussionPerm && deleteDiscussionCommentPerm {
					return key, nil
				}
			}
			return key, errors.New("community manager role undefined in the organization")
		},
	}

	for _, check := range checks {
		violationKey, err := check(o)
		if err != nil {
			violations[violationKey] = err.Error()
		}
		allErrors = errors.Join(allErrors, err)
	}

	if allErrors != nil {
		return types.Failed, &types.CheckError{
			Err:        allErrors,
			Check:      types.GoCGuardrails,
			Violations: violations,
		}
	}
	return types.Passed, nil
}
