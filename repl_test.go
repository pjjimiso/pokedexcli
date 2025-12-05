package main

import (
    "testing"
)

func TestCleanInput(t *testing.T) {
    
    cases := []struct {
        input       string
        expected    []string
    }{
        {
            input:  " hello world ",
            expected: []string{"hello", "world"},
        },
        // add more cases here
    }

    for _, c := range cases {
        actual := cleanInput(c.input)

        if len(actual) != len(c.expected) {
            t.Errorf("actual and expected slice lengths do not match")
        }

        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord { 
                t.Errorf("actual word does not equal expected word")
            }
        }
    }
}

/*
func TestGetInput(t *testing.T) { 
    cases := []struct { 
        input       string
        expected    []string
    }
}
*/
