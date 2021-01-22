// Author: xufei
// Date: 2020/11/9

package v242

/*
242. 有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true

说明:
你可以假设字符串只包含小写字母。

解题思路：

假设字母只有小写，那么可以创建一个长度26的数组，然后s字符串每个字母对应数组加一，t字符串减一，最后判断是否数组是否全部为0，是则为true。
*/

func isAnagram(s string, t string) bool {
	num := make([]int, 26)
	for _, k := range s {
		num[k-'a']++
	}
	for _, k := range t {
		num[k-'a']--
	}

	for i := 0; i < 26; i++ {
		if num[i] != 0 {
			return false
		}
	}

	return true
}
