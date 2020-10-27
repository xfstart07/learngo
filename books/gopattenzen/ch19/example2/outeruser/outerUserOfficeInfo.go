// Author: xufei
// Date: 2019-05-15

package outeruser

type IOuterUserOfficeInfo interface {
	GetUserOfficeInfo() map[string]string
}

type OuterUserOfficeInfo struct {
}

func (*OuterUserOfficeInfo) GetUserOfficeInfo() map[string]string {
	info := make(map[string]string)
	info["jobPosition"] = "这个人的职位是BOSS..."
	info["officeTelNumber"] = "员工的办公电话是"
	return info
}

func NewOuterUserOfficeInfo() *OuterUserOfficeInfo {
	return &OuterUserOfficeInfo{}
}
