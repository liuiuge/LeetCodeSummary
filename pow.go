func myPow(x float64, n int) float64 {
	if n == 0 {return 1}
    if n < 0 {
        return 1 / (x * (myPow(x, - n - 1)))
    }
    res := 1.0
    for n > 1 {
        if n % 2 == 1 {
            res *= x
        }
        x = x * x
        n /= 2
    }
    res *= x
    return res
}
