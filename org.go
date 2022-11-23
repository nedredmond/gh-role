package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/api"
)

type LoginQuery struct {
	Viewer struct {
		Login string
	}
}

func orgRole(org string, team string) (orgRole string) {
	return getOrgRole(org, team, getLogin())
}

func getOrgRole(org string, team string, login string) string {
	restClient, err := gh.RESTClient(nil)
	if err != nil {
		log.Fatal(err)
	}

	// build query
	query := []string{"orgs", org}
	if team != "" {
		query = append(query, "teams", team)
	}
	query = append(query, "memberships", login)

	path := strings.Join(query, "/")
	var resp map[string]interface{}
	err = restClient.Get(path, &resp)
	if err != nil {
		if err.(api.HTTPError).StatusCode == 404 {
			log.Fatal(fmt.Errorf("user has no role in %s", orgEntity(org, team)))
		}
		log.Fatal(err)
	}

	return resp["role"].(string)
}

func orgEntity(org string, team string) string {
	if team != "" {
		return team
	}
	return org
}

func getLogin() string {
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

	return query.Viewer.Login
}
