// Author: Xu Fei
// Date: 2018/8/24
package word

import "testing"

func Benchmark_IsPalindrome(b *testing.B) {
	for i := 0; i<b.N; i++ {
		IsPalindrome("A man, a Plan, a canal: Panama")
	}
}
