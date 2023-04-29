package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Flags for the command
	repo := flag.String("r", "", "The repo for which to check roles. If blank, the current repo is used.")
	org := flag.String("o", "", "The org for which to check roles. If blank, defaults to repo check. If present, repo flag will be ignored.")
	team := flag.String("t", "", "The team for which to check roles. Only valid in combination with org flag.")
	user := flag.String("u", "", "The user login for which to check roles. If blank, the current user is used.")
	host := flag.String("hostname", "", "The host for which to check roles. If blank, defaults to the gh config. Note that you will need to be be authenticated for the host through the gh cli.")
	friendly := flag.Bool("f", false, "Prints a friendly message. Otherwise, prints a machine-readable role name.")
	// Overrides default help message to inform about args
	defaultUsage := flag.Usage
	flag.Usage = func() {
		defaultUsage()
		fmt.Println("  List roles as space-separated arguments after any other flags to check if the current user has one of those roles.")
		fmt.Println("  Will exit with a non-zero status if the user does not have one of the specified roles.")
	}
	flag.Parse()
	var roles = flag.Args()

	isViewer := *user == ""
	if isViewer {
		user = _viewerLogin()
	}

	if *org != "" {
		err := Evaluate(*user, NestedEntityName(*org, *team), OrgRole(*org, *team, *user, *host), roles, *friendly)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	// Check repo roles
	repository := _getRepo(*repo, *host)
	var userRole string
	if isViewer {
		// Much faster to use viewerPermission check
		userRole = RepoRoleForViewer(repository)
	} else {
		userRole = RepoRoleForUser(repository, *user, *host)
	}
	err := Evaluate(*user, _repoPath(repository), userRole, roles, *friendly)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
