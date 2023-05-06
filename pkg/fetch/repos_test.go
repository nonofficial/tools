package fetch

import (
	"os"
	"testing"
)

func TestFetchRepos(t *testing.T) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		t.Skip("Set GITHUB_TOKEN to run this test.")
	}
	repos, err := FetchRepos("nonofficial")
	if err != nil {
		t.Fatalf("%s", err)
	}
	for _, repo := range repos {
		t.Logf("%s", repo.GetCloneURL())
	}
}
