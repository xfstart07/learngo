// Author: xufei
// Date: 2019-05-15

package userinfo

import "fmt"

type IUserInfo interface {
	GetUserName() string
	GetMobileNumber() string
	GetHomeAddress() string
	GetHomeTelNumber() string
	GetOfficeTelNumber() string
	GetJobPosition() string
}

type UserInfo struct {
}

func NewUserInfo() *UserInfo {
	return &UserInfo{}
}

func (*UserInfo) GetUserName() string {
	fmt.Println("名字叫做...")
	return ""
}

func (*UserInfo) GetMobileNumber() string {
	fmt.Println("手机号叫做...")
	return ""
}

func (*UserInfo) GetHomeAddress() string {
	fmt.Println("家庭地址叫做...")
	return ""
}

func (*UserInfo) GetHomeTelNumber() string {
	fmt.Println("家庭电话叫做...")
	return ""
}

func (*UserInfo) GetOfficeTelNumber() string {
	fmt.Println("办公电话叫做...")
	return ""
}

func (*UserInfo) GetJobPosition() string {
	fmt.Println("职位叫做...")
	return ""
}
