package repositoryset

import (
	"gh_foundations/internal/pkg/functions"
	"log"
	"os"

	"github.com/tidwall/gjson"

	githubfoundations "gh_foundations/internal/pkg/types/github_foundations"
)

func genFromTerraformerFile(stateFile string) *githubfoundations.RepositorySetInput {
	stateBytes, err := os.ReadFile(stateFile)
	if err != nil {
		log.Fatalf("Error reading state file %s. %s", stateFile, err.Error())
	}
	result := gjson.Parse(string(stateBytes))

	list := result.Get("modules.0.resources").Map()
	repositorySets := new(githubfoundations.RepositorySetInput)
	repositoryUserPermissions := make(map[string]map[string]string)
	for resource_id, gjsonResult := range list {
		rType := functions.IdentifyFoundationsResourceType(resource_id)
		rAttributes := gjsonResult.Get("primary.attributes")
		if rType == githubfoundations.Repository {
			repository := functions.MapTerraformerRepositoryToGithubFoundationRepository(rAttributes)
			visibility := rAttributes.Get("visibility").String()
			if visibility == "public" {
				repositorySets.PublicRepositories = append(repositorySets.PublicRepositories, repository)
			} else {
				repositorySets.PrivateRepositories = append(repositorySets.PrivateRepositories, repository)
			}
		} else if rType == githubfoundations.RepositoryCollaborator {
			repositoryName := rAttributes.Get("repository").String()
			permission := rAttributes.Get("permission").String()
			username := rAttributes.Get("username").String()
			userPermission, ok := repositoryUserPermissions[repositoryName]
			if !ok {
				userPermission = make(map[string]string)
			}
			userPermission[username] = permission
			repositoryUserPermissions[repositoryName] = userPermission
		}
	}

	for _, repository := range repositorySets.PrivateRepositories {
		repository.UserPermissions = repositoryUserPermissions[repository.Name]
	}
	for _, repository := range repositorySets.PublicRepositories {
		repository.UserPermissions = repositoryUserPermissions[repository.Name]
	}

	return repositorySets
}
