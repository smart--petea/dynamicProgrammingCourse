package fcc

import (
    "testing"
)

func TestFib(t *testing.T) {
    testEntries := []TestEntry{
        {
            Input: 1,
            Expected: 1,
        },
        {
            Input: 2,
            Expected: 1,
        },
        {
            Input: 6,
            Expected: 8,
        },
        {
            Input: 7,
            Expected: 13,
        },
        {
            Input: 8,
            Expected: 21,
        },
        {
            Input: 50,
            Expected: 12586269025,
        },
    }

    for _, testEntry := range testEntries {
        output := Fib(testEntry.Input)
        if output != testEntry.Expected {
            t.Errorf("Input=%d, Output=%d, Expected=%d\n", testEntry.Input, output, testEntry.Expected)
        }
    }
}
