// Author: xufei
// Date: 2020/3/24

package v292

// 巴什博弈
// 题解：如果n=m+1，那么由于一次最多只能取m个，所以，无论先取者拿走多少个，后取者都能够一次拿走剩余的物品，后者取胜。
// 所以要保持给对手留下 (m+1) 的倍数，就能最后获胜
func canWinNim(n int) bool {
	m := 4
	return !(n%m == 0)
}
