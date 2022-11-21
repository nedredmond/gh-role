package main

import (
	"flag"
	"fmt"
)

func main() {
	// Flags for the command
	repo := flag.String("r", "", "The repo for which to check roles. If blank, the current repo is used.")
	org := flag.String("o", "", "The org for which to check roles. If blank, defaults to repo check. If present, `-r` flag will be ignored.")
	friendly := flag.Bool("f", false, "Prints a friendly message. Otherwise, prints a machine-readable string.")
	// Overrides default help message to inform about args
	defaultUsage := flag.Usage
	flag.Usage = func() {
		defaultUsage()
		fmt.Println("  List roles as space-separated arguments after any other flags to check if the current user has one of those roles.")
		fmt.Println("  Will exit with a non-zero status if the user does not have one of the specified roles.")
	}
	flag.Parse()
	var roles = flag.Args()

	if *org != "" {
		// Check org roles
		evaluate(*org, orgRole(*org), roles, *friendly)
	}

	// Check repo roles
	if *repo == "" {
		repo = currentRepoName()
	}
	evaluate(*repo, repoRole(*repo), roles, *friendly)
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
