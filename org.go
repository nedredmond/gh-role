package main

import (
	"fmt"
	"log"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/api"
)

type LoginQuery struct {
	Viewer struct {
		Login string
	}
}

func orgRole(org string) (orgRole string) {
	return getOrgRole(org, getLogin())
}

func getOrgRole(org string, login string) string {
	restClient, err := gh.RESTClient(nil)
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("orgs/%s/memberships/%s", org, login)
	var resp map[string]interface{}
	err = restClient.Get(path, &resp)
	if err != nil {
		if err.(api.HTTPError).StatusCode == 404 {
			log.Fatal(fmt.Errorf("user has no role in %s", org))
		}
		log.Fatal(err)
	}

	return resp["role"].(string)
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
