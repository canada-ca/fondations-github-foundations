package status

import (
	githubfoundations "gh_foundations/internal/pkg/types/github_foundations"
)


type Inputs struct {
	DefaultRepositoryTeamPermissions 	map[string]string		`mapstructure:"default_repository_team_permissions"`
	PrivateRepositories					map[string]Repository	`mapstructure:"private_repositories"`
	PublicRepositories 					map[string]Repository	`mapstructure:"public_repositories"`
}

type Repository struct {
	Name 						string		`mapstructure:",label"`
	AllowUpdateBranch 			bool 		`mapstructure:"allow_update_branch"`
	AdvanceSecurity 			bool 		`mapstructure:"advance_security"`
	AllowAutoMerge 				bool 		`mapstructure:"allow_auto_merge"`
	DefaultBranch 				string 		`mapstructure:"default_branch"`
	DeleteHeadBranchOnMerge		bool 		`mapstructure:"delete_head_on_merge"`
	DependabotSecurityUpdates	bool 		`mapstructure:"dependabot_security_updates"`
	Description 				string 		`mapstructure:"description"`
	HasVulnerabilityAlerts 		bool 		`mapstructure:"has_vulnerability_alerts"`
	Homepage 					string 		`mapstructure:"homepage"`
	ProtectedBranches 			[]string	`mapstructure:"protected_branches"`
	RequiresWebCommitSignOff 	bool 		`mapstructure:"requires_web_commit_signing"`
	Topics 						[]string	`mapstructure:"topics"`
}


type OrgProjectSet struct {
	RepositorySets 		map[string]githubfoundations.RepositorySetInput
}

type OrgSet struct {
	OrgProjectSets 	map[string]OrgProjectSet
}

// Return only the names of the repositories managed by the tool
// that have GHAS enabled
func (org OrgSet) WithGHASEnabled() OrgSet {
	reposWithGHAS := OrgSet{}
	reposWithGHAS.OrgProjectSets = make(map[string]OrgProjectSet)

	for orgName, projects := range org.OrgProjectSets {
		ptrOrgProjectSet := new(OrgProjectSet)
		ptrOrgProjectSet.RepositorySets = make(map[string]githubfoundations.RepositorySetInput)
		reposWithGHAS.OrgProjectSets[orgName] = *ptrOrgProjectSet

		for projectName, repoSet := range projects.RepositorySets {
			ptrRepositorySetInput := new(githubfoundations.RepositorySetInput)

			ptrRepositorySetInput.DefaultRepositoryTeamPermissions = repoSet.DefaultRepositoryTeamPermissions
			ptrRepositorySetInput.PublicRepositories = repoSet.PublicRepositories
			for _, repo := range repoSet.PrivateRepositories {
				if repo.AdvanceSecurity {
					ptrRepositorySetInput.PrivateRepositories = append(ptrRepositorySetInput.PrivateRepositories, repo)
				}
			}
			reposWithGHAS.OrgProjectSets[orgName].RepositorySets[projectName] = *ptrRepositorySetInput

		}
	}
	return reposWithGHAS
}


// Given a repository struct returned by the HCL parser, return a githubfoundations.RepositoryInput
func (repo *Repository) GetRepositoryInput() githubfoundations.RepositoryInput {
	return githubfoundations.RepositoryInput{
		Name: repo.Name,
		AdvanceSecurity: repo.AdvanceSecurity,
		AllowAutoMerge: repo.AllowAutoMerge,
		DefaultBranch: repo.DefaultBranch,
		DeleteHeadBranchOnMerge: repo.DeleteHeadBranchOnMerge,
		DependabotSecurityUpdates: repo.DependabotSecurityUpdates,
		Description: repo.Description,
		HasVulnerabilityAlerts: repo.HasVulnerabilityAlerts,
		Homepage: repo.Homepage,
		ProtectedBranches: repo.ProtectedBranches,
		RequiresWebCommitSignOff: repo.RequiresWebCommitSignOff,
		Topics: repo.Topics,
	}
}
