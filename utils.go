package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func evaluate(entity string, userRole string, roles []string, friendly bool) {
	// If no roles were specified, print the role
	if len(roles) == 0 {
		succeed(entity, userRole, friendly)
	}

	// Otherwise, check if the user has one of the specified roles
	for _, role := range roles {
		if strings.EqualFold(userRole, role) {
			succeed(entity, userRole, friendly)
		}
	}

	// If we got here, the user doesn't have any of the specified roles
	fail(entity, userRole, roles)
}

func succeed(entity string, userRole string, friendly bool) {
	roleString := strings.ToLower(userRole)
	if friendly {
		fmt.Printf("User has %s role in %s.\n", roleString, entity)
	} else {
		fmt.Println(roleString)
	}
	os.Exit(0)
}

func fail(entity string, userRole string, checkedRoles []string) {
	s := ""
	if len(checkedRoles) > 1 {
		s = "s"
	}

	log.Fatal(
		fmt.Errorf(
			"user does not have role%s in %s: %s; found %s",
			s, entity, strings.Join(checkedRoles, ", "), strings.ToLower(userRole),
		),
	)
}
