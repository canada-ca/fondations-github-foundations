package github

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/go-github/v61/github"
)

type IGithubService interface {
	GetOrganization(slug string) (Organization, error)
	GetRepositories(owner string, filterFn func(r Repository) bool) ([]Repository, error)
}

type GithubService struct {
	client *github.Client
}

func NewGithubService(authToken string) IGithubService {
	return &GithubService{
		client: github.NewClient(nil).WithAuthToken(authToken),
	}
}

func (g *GithubService) GetOrganization(slug string) (Organization, error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancelFn()
	o, _, err := g.client.Organizations.Get(ctx, slug)
	if err != nil {
		return Organization{}, err
	}

	return Organization{
		Organization: o,
	}, nil
}

func (g *GithubService) GetRepositories(owner string, filterFn func(r Repository) bool) ([]Repository, error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancelFn()

	repos, _, err := g.client.Repositories.ListByUser(ctx, owner, nil)
	if err != nil {
		return []Repository{}, err
	}

	var repositories []Repository

	for _, r := range repos {
		rules, _, err := g.client.Repositories.GetRulesForBranch(ctx, owner, r.GetName(), r.GetDefaultBranch())
		var rulesets []map[string]interface{}
		if err == nil {
			rulesetBytes, err := json.Marshal(rules)
			if err == nil {
				json.Unmarshal(rulesetBytes, &rulesets)
			}
		}

		repositories = append(repositories, Repository{
			slug:       r.GetName(),
			rulesets:   rulesets,
			Repository: r,
		})
	}

	return repositories, nil
}
