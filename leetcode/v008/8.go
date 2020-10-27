package v008

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// 边界情况:
// 1. 为空，或只有空白符
// 2. 数字超出11长度，或值太大，太小
// 3. 数字字符串前面很多个0，正数，负数的情况
// 考察点：边界情况比较多
func myAtoi(str string) int {
	re := regexp.MustCompile(`^(\s*)?(\-|\+)?\d+`)
	digStr := re.FindString(strings.Trim(str, " "))

	// 消除前缀为0
	digStr = strings.TrimLeft(digStr, "0")

	if len(digStr) == 0 {
		return 0
	}

	if digStr[0] == '-' {
		digStr = string(digStr[0]) + strings.TrimLeft(digStr[1:], "0")
	}

	// fmt.Println(digStr)

	// 超出长度范围
	if len(digStr) > 11 {
		if digStr[0] == '-' {
			return -2147483648
		} else {
			return 2147483647
		}
	}

	dig, err := strconv.Atoi(digStr)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	max := 2147483647
	min := -2147483648

	if dig > max {
		return 2147483647
	}

	if dig < min {
		return -2147483648
	}

	return dig
}
