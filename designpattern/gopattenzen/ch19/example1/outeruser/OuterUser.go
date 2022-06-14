// Author: xufei
// Date: 2019-05-15

package outeruser

type IOuterUser interface {
	GetUserBaseInfo() map[string]string
	GetUserHomeInfo() map[string]string
	GetUserOfficeInfo() map[string]string
}

type OuterUser struct {
}

func NewOuterUser() *OuterUser {
	return &OuterUser{}
}

func (*OuterUser) GetUserBaseInfo() map[string]string {
	baseInfo := make(map[string]string)
	baseInfo["userName"] = "这个员工叫混世魔王..."
	baseInfo["mobileNumber"] = "这个员工手机..."
	return baseInfo
}

func (*OuterUser) GetUserHomeInfo() map[string]string {
	homeInfo := make(map[string]string)
	homeInfo["homeAddress"] = "这个员工家底..."
	homeInfo["homeTelNumber"] = "这个员工家电话..."
	return homeInfo
}

func (*OuterUser) GetUserOfficeInfo() map[string]string {
	officeInfo := make(map[string]string)
	officeInfo["jobPosition"] = "这个员工办公室..."
	officeInfo["officeTelNumber"] = "这个员工办公室..."
	return officeInfo
}
