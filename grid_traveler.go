package fcc

func GridTraveler(m, n int) int {
    if m == 1 && n == 1 {
        return 1
    }

    if m == 0 || n == 0 { 
        return 0
    }

    return GridTraveler(m, n-1) + GridTraveler(m-1, n)
}
