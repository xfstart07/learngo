package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestExampleRegexNumber(t *testing.T) {
	ExampleRegNumber()
}

func TestPatternOr(t *testing.T) {
	//  xy 匹配x y
	// x|y  匹配x或者y 优先x
	source := "asdfdsxxxyyfergsfasfxyfa"
	pattern := `x|y|a`
	reg := regexp.MustCompile(pattern)
	ret := reg.FindStringIndex(source)
	fmt.Println(ret)
}

func TestPatternN2M(t *testing.T) {
	//x* 匹配零个或者多个x,优先匹配多个
	//x+ 匹配一个或者多个x，优先匹配多个
	//x? 匹配零个或者一个x，优先匹配一个
	//source = "xxxxewexxxasdfdsxxxyyfergsfasfxyfa"
	//pattern = `x*`

	// x{n,m} 匹配n个到m个x，优先匹配m个
	// x{n,}  匹配n个到多个x，优先匹配更多
	// x{n} 或者x{n}?  只匹配n个x
	source = "xxxxxxxewexxxasdfdsxxxyyfergsfasfxyfa"
	pattern = `x{4,}`

	reg := regexp.MustCompile(pattern)
	ret := reg.FindString(source)
	fmt.Println(ret)
}

func TestPatternPriorityN(t *testing.T) {
	// x{n,m}? 匹配n个到m个x，优先匹配n个
	// x{n,}?  匹配n个到多个x，优先匹配n个
	// x*?   匹配零个或者多个x，优先匹配0个
	// x+?   匹配一个或者多个x，优先匹配1个
	// x??   匹配零个或者一个x，优先匹配0个
	source = "xxxxxxxewexxxasdfdsxxxyyfergsfasfxyfa"
	pattern = `x+?`

	reg := regexp.MustCompile(pattern)
	ret := reg.FindString(source)
	fmt.Println(ret)
}

func TestPatternNumberN2M(t *testing.T) {
	//[\d] 或者[^\D] 匹配数字
	//[^\d]或者 [\D] 匹配非数字
	source = "xx435ff5237yy6346fergsfasfxyfa"
	pattern = `[\d]{3,}` //匹配3个或者更多个数字

	reg := regexp.MustCompile(pattern)
	ret := reg.FindString(source)
	fmt.Println(ret)
}

func TestPatternAlpha(t *testing.T) {
	source = "xx435ffGUTEYgjk52RYPHFY37yy6346ferg6987sfasfxyfa"
	pattern = `[[:alpha:]]{5,}` //5个或者多个字母，相当于A-Za-z

	reg := regexp.MustCompile(pattern)
	ret := reg.FindString(source)
	fmt.Println(ret)
}

func TestPatternHan(t *testing.T) {
	source = "xx435,./$%(*(_&jgshgs发个^$%ffG返回福hjh放假啊啥UTEYgjk52RYPHFY37yy6346ferg6987sfasfxyfa"
	pattern = `[\p{Han}]+` //匹配连续的汉字

	reg := regexp.MustCompile(pattern)
	ret := reg.FindString(source)
	fmt.Println(ret)
}

func TestPatternPhone(t *testing.T) {
	source = "13244820821HG74892109977HJA15200806084S11233240697hdgsfhah假发发货"
	pattern = `1[3|5|7|8|9][\d]{9}` //匹配电话号码

	reg := regexp.MustCompile(pattern)
	ret := reg.FindString(source)
	fmt.Println(ret)
}

func TestPatternEmail(t *testing.T) {
	source = "132@12.comGKGk15@163.cn200806084S11233240697hdgsfhah假发发货"
	pattern = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱

	reg := regexp.MustCompile(pattern)
	ret := reg.FindString(source)
	fmt.Println(ret)
}

func TestPatternOther(t *testing.T) {
	//匹配用户名或者密码 `^[a-zA-Z0-9_-]{4,16}$`  字母或者数字开头，区分大小写，最短4位最长16位
	//匹配IP地址1 `^$(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`

	//匹配IP地址2
	//pattern = `((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)`

	// 匹配IP地址3
	// pattern = `(\d+).(\d+).(\d+).(\d+)`

	// 匹配IP地址4

	//匹配日期 年-月-日 `(\d{4}|\d{2})-((1[0-2])|(0?[1-9]))-(([12][0-9])|(3[01])|(0?[1-9]))`
	//匹配日期 月-日-年  `((1[0-2])|(0?[1-9]))/(([12][0-9])|(3[01])|(0?[1-9]))/(\d{4}|\d{2})`
	//匹配时间 小时：分钟 24小时制 ` ((1|0?)[0-9]|2[0-3]):([0-5][0-9]) `
	//匹配邮编  `[1-9][\d]5`
	//匹配URL `[a-zA-z]+://[^\s]*`
}

func ExampleRegNumber() {
	text := "132121111"
	reg := regexp.MustCompile(`\d{4}$`)
	regText := reg.ReplaceAllString(text, "0000")
	fmt.Println(regText)
}
