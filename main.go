package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cli/go-gh"
)

func main() {
	// Flags for the command
	repo := flag.String("r", "", "The repo for which to check roles. If blank, the current repo is used.")
	friendly := flag.Bool("f", false, "Prints a friendly message instead of the role constant.")
	// Overrides default help message to inform about args
	defaultUsage := flag.Usage
	flag.Usage = func() {
		defaultUsage()
		fmt.Println("  List roles as space-separated arguments after any other flags to check if the current user has one of those roles.")
		fmt.Println("  Will exit with a non-zero status if the user does not have one of the specified roles.")
	}
	flag.Parse()
	var roles = flag.Args()

	// Start building the command
	ghArgs := []string{"repo", "view"}
	if repo != nil && *repo != "" {
		// If repo provided, add it to the command
		ghArgs = append(ghArgs, *repo)
	} else {
		// Write the current repo to the variable but we don't add it to the command
		// It's for later
		repo = currentRepoName()
	}
	ghArgs = append(ghArgs, "--json", "viewerPermission")

	// Execute the command
	stdOut, _, err := gh.Exec(ghArgs...)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Parse the output
	var result struct {
		Role string `json:"viewerPermission"`
	}
	err = json.Unmarshal(stdOut.Bytes(), &result)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	success := func() {
		if *friendly {
			fmt.Printf("Current user has %s role on %s.\n", result.Role, *repo)
		} else {
			fmt.Println(result.Role)
		}
	}

	// If no roles were specified, print the role
	if len(roles) == 0 {
		success()
		os.Exit(0)
	}
	// Otherwise, check if the user has one of the specified roles
	for _, role := range roles {
		if strings.EqualFold(result.Role, role) {
			success()
			os.Exit(0)
		}
	}
	// If we got here, the user doesn't have any of the specified roles

	s := ""
	if len(roles) > 1 {
		s = "s"
	}

	log.Fatal(
		fmt.Errorf(
			"User does not have role%s on %s: %s; found %s",
			s, *repo, strings.Join(roles, ", "), result.Role,
		),
	)
	os.Exit(1)
}

func currentRepoName() *string {
	repository, err := gh.CurrentRepository()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	repoName := repository.Name()
	return &repoName
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
