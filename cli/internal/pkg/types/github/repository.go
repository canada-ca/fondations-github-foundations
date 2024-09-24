package github

import (
	"errors"
	"fmt"
	"gh_foundations/internal/pkg/types"
	"reflect"
	"time"

	"github.com/google/go-github/v61/github"
)

type Repository struct {
	slug     string
	rulesets []map[string]any
	*github.Repository
}

func (r *Repository) Check(checkTypes []types.CheckType) types.CheckReport {
	report := types.CheckReport{
		EntityType: "github_repository",
		EntityId:   r.slug,
		Timestamp:	time.Now().Format(time.RFC3339), // Syslog compliant timestamp
		Results:    make(map[types.CheckType]types.CheckResult),
		Errors:     []types.CheckError{},
	}
	for _, t := range checkTypes {
		switch t {
		case types.GoCGuardrails:
			r, err := r.GoCGuardrailsCompliant()
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

func (r *Repository) GoCGuardrailsCompliant() (types.CheckResult, *types.CheckError) {
	var allErrors error
	violations := make(map[string]string)
	checks := []func(repo *Repository) (string, error) {
		func(repo *Repository) (string, error) {
			key := "dependabot_security_updates"
			if repo.GetSecurityAndAnalysis().GetDependabotSecurityUpdates().GetStatus() != "enabled" {
				return key, errors.New("dependabot_security_updates is not enabled. Expected it to be enabled")
			}
			return key, nil
		},
		func(repo *Repository) (string, error) {
			key := "secret_scanning"
			if repo.GetSecurityAndAnalysis().GetSecretScanning().GetStatus() != "enabled" {
				return key, errors.New("secret_scanning is not enabled. Expected it to be enabled")
			}
			return key, nil
		},
		func(repo *Repository) (string, error) {
			key := "secret_scanning_push_protection"
			if repo.GetSecurityAndAnalysis().GetSecretScanningPushProtection().GetStatus() != "enabled" {
				return key, errors.New("secret_scanning_push_protection is not enabled. Expected it to be enabled")
			}
			return key, nil
		},
		func(repo *Repository) (string, error) {
			key := "delete_branch_on_merge"
			if !repo.GetDeleteBranchOnMerge() {
				return key, errors.New("delete_branch_on_merge is not enabled. Expected it to be enabled")
			}
			return key, nil
		},
		func(repo *Repository) (string, error) {
			key := "rulesets"
			var errs error
			expectedRulesets := []map[string]any{
				{
					"type": "pull_request",
					"parameters": map[string]any{
						"required_approving_review_count":   1,
						"dismiss_stale_reviews_on_push":     true,
						"require_code_owner_review":         false,
						"require_last_push_approval":        true,
						"required_review_thread_resolution": false,
					},
					"ruleset_source_type": "Repository",
				},
			}

			for _, ruleset := range expectedRulesets {
				eq := reflect.DeepEqual(ruleset, r.rulesets)
				if !eq {
					errMsg := fmt.Sprintf("expected ruleset %q to be %q. Got %q instead", ruleset, fmt.Sprint(ruleset), fmt.Sprint(r.rulesets))
					violations["rulesets"] = errMsg
					errs = errors.Join(errs, errors.New(errMsg))
				}
			}
			return key, errs
		},
	}

	for _, check := range checks {
		violationKey, err := check(r)
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
