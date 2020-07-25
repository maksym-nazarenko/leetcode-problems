package tasks

func countPrimes(n int) int {

	var isPrime []bool = make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}

	for i := 2; i*i < n; i++ {
		if !isPrime[i] {
			continue
		}

		for j := i * i; j < n; j += i {
			isPrime[j] = false
		}
	}

	ret := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			ret++
		}
	}

	return ret
}
