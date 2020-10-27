// Author: xufei
// Date: 2019-05-15

package outeruser

type IOuterUserBaseInfo interface {
	GetUserBaseInfo() map[string]string
}

type OuterUserBaseInfo struct {
}

func (*OuterUserBaseInfo) GetUserBaseInfo() map[string]string {
	baseInfo := make(map[string]string)
	baseInfo["userName"] = "这个员工叫混世魔王..."
	baseInfo["mobileNumber"] = "这个员工电话是"
	return baseInfo
}

func NewOuterUserBaseInfo() *OuterUserBaseInfo {
	return &OuterUserBaseInfo{}
}
