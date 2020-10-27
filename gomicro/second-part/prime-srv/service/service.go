// Author: xufei
// Date: 2020/3/5

package service

func GetPrime(n int64) []int64 {
	var outputs []int64

	var i int64
	for i = 1; i <= n; i++ {
		if isPrime(i) {
			outputs = append(outputs, i)
		}
	}

	return outputs
}

func isPrime(i int64) bool {
	if i <= 3 {
		return true
	}

	var j int64
	for j = 2; j <= i/2; j++ {
		if i%j == 0 {
			return false
		}
	}

	return true
}
