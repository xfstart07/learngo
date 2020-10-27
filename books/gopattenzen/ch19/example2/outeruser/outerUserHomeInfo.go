// Author: xufei
// Date: 2019-05-15

package outeruser

type IOuterUserHomeInfo interface {
	GetUserHomeInfo() map[string]string
}

type OuterUserHomeInfo struct {
}

func (*OuterUserHomeInfo) GetUserHomeInfo() map[string]string {
	info := make(map[string]string)
	info["homeAddress"] = "员工的家庭地址是..."
	info["homeTelNumber"] = "员工的家庭电话是"
	return info
}

func NewOuterUserHomeInfo() *OuterUserHomeInfo {
	return &OuterUserHomeInfo{}
}
