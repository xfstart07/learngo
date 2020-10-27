// Author: xufei
// Date: 2019-05-15

package outeruser

type OuterUserInfo struct {
	OuterUser

	BaseInfo   map[string]string
	HomeInfo   map[string]string
	OfficeInfo map[string]string
}

func NewOuterUserInfo(outerUser OuterUser) *OuterUserInfo {
	return &OuterUserInfo{
		OuterUser:  outerUser,
		BaseInfo:   outerUser.GetUserBaseInfo(),
		HomeInfo:   outerUser.GetUserHomeInfo(),
		OfficeInfo: outerUser.GetUserOfficeInfo(),
	}
}

func (o *OuterUserInfo) GetUserName() string {
	return o.BaseInfo["userName"]
}

func (o *OuterUserInfo) GetMobileNumber() string {
	return o.BaseInfo["userName"]
}

func (o *OuterUserInfo) GetHomeAddress() string {
	return o.HomeInfo["homeAddress"]
}

func (o *OuterUserInfo) GetHomeTelNumber() string {
	return o.HomeInfo["homeTelNumber"]
}

func (o *OuterUserInfo) GetOfficeTelNumber() string {
	return o.OfficeInfo["officeTelNumber"]
}

func (o *OuterUserInfo) GetJobPosition() string {
	return o.OfficeInfo["jobPosition"]
}
