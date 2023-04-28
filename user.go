package main

import (
	"log"

	"github.com/cli/go-gh"
)

func _viewerLogin() *string {
	// GitHub CLI doesn't make it easy to get the current user's login.
	// I could either parse it from the verbose `gh auth status` or make an API call.
	// Since the status is subject to change, I'll just make the API call.
	gqlClient, err := gh.GQLClient(nil)
	if err != nil {
		log.Fatal(err)
	}

	var query LoginQuery
	err = gqlClient.Query("LoginQuery", &query, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &query.Viewer.Login
}
