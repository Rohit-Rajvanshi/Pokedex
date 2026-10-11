package main

import (
	"testing"	
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string 
	}{
		{
			input : "   hello " ,
			expected : []string{"hello"},
		},
		{
			input : "Hello             World ",
			expected : []string{"hello", "world"},
		},
		{
			input : "This world shall Perish",
			expected : []string{"this", "world" , "shall" , "perish"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
	
		if len(actual) != len(c.expected) {
			t.Errorf("Failed the test",)
			continue 
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if (word != expectedWord) {
				t.Errorf("Failed the test")
			}

		}
	}
}