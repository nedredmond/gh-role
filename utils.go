package main

import (
	"log"
)

func NestedEntityName(entities ...string) string {
	for i := range entities {
		// Iterate backwards over list of entities
		j := len(entities) - 1 - i
		if entities[j] != "" {
			return entities[j]
		}
	}
	panic("no named entities")
}

func noRoleErr(user string, entity string) {
	log.Fatalf("%s has no role in %s", user, entity)
}
