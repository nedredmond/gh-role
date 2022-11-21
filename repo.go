package main

import (
	"encoding/json"
	"log"

	"github.com/cli/go-gh"
)

func currentRepoName() *string {
	repository, err := gh.CurrentRepository()
	if err != nil {
		log.Fatal(err)
	}
	repoName := repository.Name()
	return &repoName
}

func repoRole(repo string) (repoRole string) {
	// Build the command
	ghArgs := []string{"repo", "view", repo, "--json", "viewerPermission"}

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
