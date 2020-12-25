package fcc

import (
    "testing"
)

type TestEntryMN struct {
    InputM int
    InputN int
    Expected int
}

func TestGridTraveler(t *testing.T) {
    testEntries := []TestEntryMN{
        {
            InputM: 1,
            InputN: 1,
            Expected: 1,
        },
        {
            InputM: 2,
            InputN: 3,
            Expected: 3,
        },
        {
            InputM: 3,
            InputN: 2,
            Expected: 3,
        },
        {
            InputM: 3,
            InputN: 3,
            Expected: 6,
        },
        /*
        {
            InputM: 18,
            InputN: 18,
            Expected: 6,
        },
        */
    }

    for _, testEntry := range testEntries {
        output := GridTraveler(testEntry.InputM, testEntry.InputN)
        if output != testEntry.Expected {
            t.Errorf("(%d,%d)->%d, expected=%d\n", testEntry.InputM, testEntry.InputN, output, testEntry.Expected)
        }
    }
}
