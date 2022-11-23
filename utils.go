package main

func NestedEntityName(entities ...string) string {
	for i := range entities {
		// Iterate bacwards over list of entities
		j := len(entities) - 1 - i
		if entities[j] != "" {
			return entities[j]
		}
	}
	panic("no named entities")
}
