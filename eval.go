package main

import (
	"fmt"
	"strings"
)

func Evaluate(entity string, userRole string, roles []string, friendly bool) (err error) {
	// If no roles were specified, print the role
	if len(roles) == 0 {
		_succeed(entity, userRole, friendly)
		return
	}

	// Otherwise, check if the user has one of the specified roles
	for _, role := range roles {
		if strings.EqualFold(userRole, role) {
			_succeed(entity, userRole, friendly)
			return
		}
	}

	// If we got here, the user doesn't have any of the specified roles
	return _fail(entity, userRole, roles)
}

func _succeed(entity string, userRole string, friendly bool) {
	roleString := strings.ToLower(userRole)
	if friendly {
		fmt.Printf("User has %s role in %s.\n", roleString, entity)
	} else {
		fmt.Println(roleString)
	}
}

func _fail(entity string, userRole string, checkedRoles []string) error {
	s := ""
	if len(checkedRoles) > 1 {
		s = "s"
	}

	return fmt.Errorf(
		"user does not have role%s in %s: %s; found %s",
		s, entity, strings.Join(checkedRoles, ", "), strings.ToLower(userRole),
	)
}
