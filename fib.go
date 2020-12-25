package fcc

func Fib(n int) int {
    if n <= 2 {
        return 1
    }

    memo := make([]int, n, n)
    memo[0] = 1
    memo[1] = 1

    return fib(n, memo)
}

func fib(n int, memo []int) int {
    if memo[n-1] == 0 {
        memo[n-1] = fib(n-1, memo) + fib(n-2, memo)
    }

    return memo[n-1]
}
