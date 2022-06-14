// Author: xufei
// Date: 2019-05-15

package main

import "learngo/books/gopattenzen/ch19/example2/outeruser"

func main() {
	baseInfo := outeruser.NewOuterUserBaseInfo()
	homeInfo := outeruser.NewOuterUserHomeInfo()
	officeInfo := outeruser.NewOuterUserOfficeInfo()

	youngGirl := outeruser.NewOuterUserInfo(baseInfo, homeInfo, officeInfo)
	youngGirl.GetHomeAddress()
}
