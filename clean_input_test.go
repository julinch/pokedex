package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    " Leading and trailing ",
			expected: []string{"leading", "and", "trailing"},
		},
		{
			input:    "MULTIPLE SPACES here",
			expected: []string{"multiple", "spaces", "here"},
		},
		{
			input:    "\tTabs\tand\nnewlines\n",
			expected: []string{"tabs", "and", "newlines"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "OneWord",
			expected: []string{"oneword"},
		},

		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The length of the actual output doesn't match the exprected! Actual: %v, expected: %v", len(actual), len(c.expected))
			t.Fail()
			return
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("The word of the actual output doesn't match the exprected! Actual: %s, expected: %s", word, expectedWord)
				t.Fail()
				return
			}
		}
	}
}
