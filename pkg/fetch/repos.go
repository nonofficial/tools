package fetch

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/v52/github"
)

func FetchRepos(orgname string) ([]*github.Repository, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	ctx := context.Background()
	client := github.NewTokenClient(ctx, token)
	options := &github.RepositoryListByOrgOptions{}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), orgname, options)
	return repos, err
	if err != nil {
		return nil, err
	}
	return repos, nil
}
