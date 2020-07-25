package tasks

func countPrimes(n int) int {
	var ret = 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			ret++
		}
	}

	return ret
}

func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}
