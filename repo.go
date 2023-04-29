package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cli/go-gh/v2"
	"github.com/cli/go-gh/v2/pkg/repository"
)

func _getRepo(path string, host string) (repo repository.Repository) {
	var err error
	if path == "" {
		repo, err = repository.Current()
		if host != "" {
			repo.Host = host
		}
	} else if host != "" {
		repo, err = repository.ParseWithHost(path, host)
	} else {
		repo, err = repository.Parse(path)
	}
	if err != nil {
		log.Fatal(err)
	}
	return
}

func _repoPath(repo repository.Repository) string {
	return repo.Host + "/" + repo.Owner + "/" + repo.Name
}

func RepoRoleForViewer(repo repository.Repository) (repoRole string) {
	// Build the command
	ghArgs := []string{"repo", "view", _repoPath(repo), "--json", "viewerPermission"}

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

func RepoRoleForUser(
	repo repository.Repository,
	login string,
	host string,
) (repoRole string) {
	// Build the command
	ghArgs := []string{"api", "-H", "Accept: application/vnd.github+json", "-H", "X-GitHub-Api-Version: 2022-11-28"}
	if host != "" {
		ghArgs = append(ghArgs, "--hostname", host)
	}
	ghArgs = append(ghArgs, fmt.Sprintf("/repos/%s/%s/collaborators/%s/permission", repo.Owner, repo.Name, login))

	// Execute the command
	stdOut, _, err := gh.Exec(ghArgs...)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the output
	var result struct {
		Role string `json:"permission"`
	}
	err = json.Unmarshal(stdOut.Bytes(), &result)
	if err != nil {
		log.Fatal(err)
	}

	return result.Role
}	
