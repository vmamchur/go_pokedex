package main

import (
  "testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct{
        input string
        expected []string
    }{
        {
            input: "  hello world  ",
            expected: []string{"hello", "world"},
        },
        {
            input: " Charmander Bulbasaur PIKACHU",
            expected: []string{"charmander", "bulbasaur", "pikachu"},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("input: %q - expected length %d, got %d", c.input, len(c.expected), len(actual))
        }
        
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("input: %q, - at index %d, expected %q, got %q", c.input, i, expectedWord, word)
            }
        }
    }
}

