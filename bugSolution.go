func myFuncSolution(a []int) {
    b := make([]int, len(a))
    copy(b, a)
    for i := range b {
        b[i] = b[i] * 2
    }
    fmt.Println(b)
}
func myFuncSolution2(a []int) {
    for i := 0; i < len(a); i++ {
        a[i] = a[i] * 2
    }
    fmt.Println(a)
} 