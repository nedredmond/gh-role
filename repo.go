package main

import (
	"encoding/json"
	"log"

	"github.com/cli/go-gh/v2"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/cli/go-gh/v2/pkg/repository"
	graphql "github.com/cli/shurcooL-graphql"
)

func _getRepo(path string) (repo repository.Repository) {
	var err error
	if path == "" {
		repo, err = repository.Current()
	} else {
		repo, err = repository.Parse(path)
	}
	if err != nil {
		log.Fatal(err)
	}
	return
}

func RepoRoleForViewer(repo repository.Repository) (repoRole string) {
	// Build the command
	repoPath := repo.Owner + "/" + repo.Name
	ghArgs := []string{"repo", "view", repoPath, "--json", "viewerPermission"}

	// Execute the command
	stdOut, _, err := gh.Exec(ghArgs...)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the output
	var result struct {
		Role string `json:"viewerPermission"`
	}
	err = json.Unmarshal(stdOut.Bytes(), &result)
	if err != nil {
		log.Fatal(err)
	}

	return result.Role
}

func _getRepoPath(repo repository.Repository) string {
	return repo.Owner + "/" + repo.Name
}

func RepoRoleForUser(
	repo repository.Repository,
	login string,
) (repoRole string) {
	client, err := api.DefaultGraphQLClient()
	if err != nil {
		log.Fatal(err)
	}
	var query struct {
		Repository struct {
			Collaborators struct {
				Edges []struct {
					Permission string
				}
			} `graphql:"collaborators(login: $login)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}
	variables := map[string]interface{}{
		"owner": graphql.String(repo.Owner),
		"name":  graphql.String(repo.Name),
		"login": graphql.String(login),
	}
	err = client.Query("UserRepoPermission", &query, variables)
	if err != nil {
		log.Fatalf("unable to query %s's role in %s: %s", login, _getRepoPath(repo), err)
	}
	if err != nil || len(query.Repository.Collaborators.Edges) == 0 {
		noRoleErr(login, _getRepoPath(repo))
	}
	return query.Repository.Collaborators.Edges[0].Permission
}
