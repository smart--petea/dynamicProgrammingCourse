package fcc

type Arr2D struct {
    Arr []int
    M int
    N int
}

func (a *Arr2D) Index(i,j int) int {
    return (i - 1) * a.N + j - 1
}

func (a *Arr2D) Get(i,j int) int {
    return a.Arr[a.Index(i,j)]
}

func (a *Arr2D) Set(i,j,val int) {
    a.Arr[a.Index(i,j)] = val
}

func NewArr2D(m,n int) *Arr2D {
    return &Arr2D{
        M: m,
        N: n,
        Arr: make([]int, m*n, m*n),
    }
}
