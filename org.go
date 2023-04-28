package main

import (
	"log"
	"strings"

	"github.com/cli/go-gh/v2/pkg/api"
)

type LoginQuery struct {
	Viewer struct {
		Login string
	}
}

func OrgRole(org string, team string, login string) string {
	restClient, err := api.DefaultRESTClient()
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
		// We can't seem to parse the error response to check the status code
		// so we'll just check the error message instead.
		// https://github.com/cli/go-gh/issues/118
		if strings.Contains(err.Error(), "Not Found") {
			noRoleErr(login, NestedEntityName(org, team))
		}
		log.Fatal(err)
	}

	return resp["role"].(string)
}
