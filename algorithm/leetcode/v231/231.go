// Author: xufei
// Date: 2020/3/31

package v231

func isPowerOfTwo(n int) bool {
	f := make([]int, 40)
	f[0] = 1
	for i := 1; i < 40; i++ {
		f[i] = f[i-1] * 2
	}

	for i := 0; i < 40; i++ {
		if f[i] == n {
			return true
		}
	}

	return false
}
