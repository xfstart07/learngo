// Author: xufei
// Date: 2019-05-15

package outeruser

import "fmt"

type OuterUserInfo struct {
	BaseInfo   IOuterUserBaseInfo
	HomeInfo   IOuterUserHomeInfo
	OfficeInfo IOuterUserOfficeInfo
}

func (o OuterUserInfo) GetUserName() string {
	fmt.Println(o.BaseInfo.GetUserBaseInfo()["userName"])
	return o.BaseInfo.GetUserBaseInfo()["userName"]
}

func (o OuterUserInfo) GetMobileNumber() string {
	fmt.Println(o.BaseInfo.GetUserBaseInfo()["mobileNumber"])
	return o.BaseInfo.GetUserBaseInfo()["mobileNumber"]
}

func (o OuterUserInfo) GetHomeAddress() string {
	fmt.Println(o.HomeInfo.GetUserHomeInfo()["homeAddress"])
	return o.HomeInfo.GetUserHomeInfo()["homeAddress"]
}

func (o OuterUserInfo) GetHomeTelNumber() string {
	fmt.Println(o.HomeInfo.GetUserHomeInfo()["homeTelNumber"])
	return o.HomeInfo.GetUserHomeInfo()["homeTelNumber"]
}

func (o OuterUserInfo) GetOfficeTelNumber() string {
	fmt.Println(o.OfficeInfo.GetUserOfficeInfo()["officeTelNumber"])
	return o.OfficeInfo.GetUserOfficeInfo()["officeTelNumber"]
}

func (o OuterUserInfo) GetJobPosition() string {
	fmt.Println(o.OfficeInfo.GetUserOfficeInfo()["jobPosition"])
	return o.OfficeInfo.GetUserOfficeInfo()["jobPosition"]
}

func NewOuterUserInfo(baseInfo IOuterUserBaseInfo, homeInfo IOuterUserHomeInfo, officeInfo IOuterUserOfficeInfo) *OuterUserInfo {
	return &OuterUserInfo{
		BaseInfo:   baseInfo,
		HomeInfo:   homeInfo,
		OfficeInfo: officeInfo,
	}
}
