package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluate_fail_roles(t *testing.T) {
	assert.ErrorContains(
		t,
		Evaluate("askywalk", "council", "jedi", []string{"master", "padawan"}, true),
		"askywalk does not have roles in council: master, padawan; found jedi",
	)
}

func TestEvaluate_fail_role(t *testing.T) {
	assert.ErrorContains(
		t,
		Evaluate("askywalk", "council", "jedi", []string{"master"}, false),
		"askywalk does not have role in council: master; found jedi",
	)
}

func TestEvaluate_succeed_noError(t *testing.T) {
	assert.NoError(t, Evaluate("askywalk", "council", "jedi", []string{}, false))
}

func ExampleEvaluate_succeed_noRoles_friendly() {
	Evaluate("askywalk", "council", "jedi", []string{}, true)
	// Output: askywalk has jedi role in council.
}

func ExampleEvaluate_succeed_noRoles_unfriendly() {
	Evaluate("askywalk", "council", "jedi", []string{}, false)
	// Output: jedi
}

func ExampleEvaluate_succeed_roles_friendly() {
	Evaluate("askywalk", "council", "jedi", []string{"jedi"}, false)
	// Output: jedi
}
