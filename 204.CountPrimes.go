package code

func isPrime(n int) bool {
	for i := 3; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func CountPrimes(n int) int {
	if n > 3 {
		cnt := 2
		for i := n - 1; i > 3; i-- {
			if i%2 != 0 && isPrime(i) {
				cnt++
			}
		}
		return cnt
	} else if n < 2 {
		return 0
	} else {
		return 1
	}
}

func CountPrimesFast(n int) int {
	if n < 2 {
		return 0
	}
	flags := make([]bool, n/2)
	for i := range flags {
		flags[i] = true
	}
	for i := 3; i*i < n; i += 2 {
		if flags[i/2] {
			for j := i; j*i < n; j += 2 {
				flags[j*i/2] = false
			}
		}
	}
	c := 0
	for _, k := range flags {
		if k {
			c += 1
		}
	}
	return c
}
