package fcc

func GridTraveler(m, n int) int {
    if m == 1 && n == 1 {
        return 1
    }

    if m == 0 || n == 0 { 
        return 0
    }

    memo := NewArr2D(m, n)
    memo.Set(1, 1, 1)

    return gridTraveler(m, n, memo)
}

func gridTraveler(m, n int, memo *Arr2D) int {
    if m == 0 || n == 0 {
        return 0
    }

    val := memo.Get(m, n)
    if val == 0 {
        val = gridTraveler(m-1, n, memo) + gridTraveler(m, n-1, memo)
        memo.Set(m, n, val)
    }

    return val
}
