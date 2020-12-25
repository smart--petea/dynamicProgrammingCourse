package fcc

import (
    "testing"
)


func TestIndex(t *testing.T) {
    m := 2
    n := 3
    arr2D := NewArr2D(m,n)

    testEntries := []TestEntryMN{
        {
            InputM: 1,
            InputN: 1,
            Expected: 0,
        },
        {
            InputM: 1,
            InputN: 2,
            Expected: 1,
        },
        {
            InputM: 1,
            InputN: 3,
            Expected: 2,
        },
        {
            InputM: 2,
            InputN: 1,
            Expected: 3,
        },
        {
            InputM: 2,
            InputN: 2,
            Expected: 4,
        },
        {
            InputM: 2,
            InputN: 3,
            Expected: 5,
        },
    }

    for _, testEntry := range testEntries {
        output := arr2D.Index(testEntry.InputM, testEntry.InputN)
        if output != testEntry.Expected {
            t.Errorf("(%d,%d)->%d should point to %d", testEntry.InputM, testEntry.InputN, output, testEntry.Expected)
        }
    }
}

func TestGetSet(t *testing.T) {
    m := 5
    n := 6
    arr2D := NewArr2D(m,n)

    i := 3
    j := 4
    val := 5
    arr2D.Set(i, j, val)

    if arr2D.Get(i, j) != val {
        t.Errorf("Get is different of Set")
    }
}

func TestNewArr2D(t *testing.T) {
    m := 2
    n := 3
    arr2D := NewArr2D(m,n)

    if arr2D.M != m {
        t.Errorf("m is not set")
    }

    if arr2D.N != n {
        t.Errorf("n is not set")
    }

    if len(arr2D.Arr) != m*n {
        t.Errorf("wrong size of inner array")
    }
}
