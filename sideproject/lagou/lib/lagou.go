package lib

// 基础包
import (
	"net/http"
	"strings"

	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	. "github.com/henrylee2cn/pholcus/app/spider"           //必需
	"github.com/henrylee2cn/pholcus/common/goquery"         //DOM解析
)

//修改这个为其他岗位的，可以爬取其他岗位的数据
const positionURL = "https://www.lagou.com/zhaopin/go/?filterOption=3"

func init() {
	header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,fr;q=0.4,tr;q=0.2,zh-TW;q=0.2")
	header.Add("Connection", "keep-alive")
	//header.Add("Cookie", "user_trace_token=20170910220432-f801c133-9630-11e7-8e11-525400f775ce; LGUID=20170910220432-f801c565-9630-11e7-8e11-525400f775ce; index_location_city=%E5%85%A8%E5%9B%BD; JSESSIONID=ABAAABAAADEAAFI27EBBC4DCA6B9DBD97414B0004A32D4F; TG-TRACK-CODE=index_navigation; _gat=1; PRE_UTM=; PRE_HOST=; PRE_SITE=https%3A%2F%2Fwww.lagou.com%2Fzhaopin%2Fgo%2F3%2F%3FfilterOption%3D2; PRE_LAND=https%3A%2F%2Fwww.lagou.com%2Fzhaopin%2Fgo%2F4%2F%3FfilterOption%3D3; SEARCH_ID=418a46d847344429b67029bc1470f19c; _gid=GA1.2.1008155537.1505828050; Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1505052272,1505828050; Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1505830015; _ga=GA1.2.319466696.1505052272; LGSID=20170919220506-8a4a46e3-9d43-11e7-99b2-525400f775ce; LGRID=20170919220655-cb047879-9d43-11e7-99b2-525400f775ce")
	header.Add("Cookie", "WEBTJ-ID=20200424104142-171aa0fecfc544-0f3e5592674e5d-1c42311e-2304000-171aa0fed005cf; user_trace_token=20200424104145-3fe00d30-fa84-47fd-97c0-648d28357bfb; LGSID=20200424104145-9b624adb-143f-4796-9322-8da8f3a2cfc4; LGUID=20200424104145-da17dcb2-fff3-4b43-9577-dc0d470c3e70; _ga=GA1.2.917919990.1587696106; _gid=GA1.2.534234878.1587696107; Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1587696107; index_location_city=%E5%85%A8%E5%9B%BD; gate_login_token=4d27ce2ef03bd190a49336aa224e6af760c66df860416661; LG_LOGIN_USER_ID=e7fdebaae22c525f361c8717220b875a02331544980cd93b; LG_HAS_LOGIN=1; _putrc=E62D8C4904EFE1E6; login=true; unick=%E5%BE%90%E9%A3%9E; showExpriedIndex=1; showExpriedCompanyHome=1; showExpriedMyPublish=1; hasDeliver=4; privacyPolicyPopup=false; sajssdk_2015_cross_new_user=1; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22171aa29b3f0502-0f7c123b3d130b-1c42311e-2304000-171aa29b3f165c%22%2C%22%24device_id%22%3A%22171aa29b3f0502-0f7c123b3d130b-1c42311e-2304000-171aa29b3f165c%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_referrer%22%3A%22%22%2C%22%24latest_referrer_host%22%3A%22%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%7D%7D; _gat=1; SEARCH_ID=7961e5c8ad794174a2ae5ab14e35ec2d; X_HTTP_TOKEN=338efcf46c52693e3358967851d8083c87fd9965db; LGRID=20200424112213-5b07c119-530a-4376-b3c9-83e0847b5716; Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1587698533")
	header.Add("DNT", "1")
	header.Add("Host", "www.lagou.com")
	header.Add("Referer", "https://www.lagou.com/")
	header.Add("Upgrade-Insecure-Requests", "1")
	header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 9_3_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3112.113 Safari/535.36\"")
	lagou.Register()
}

var lagou = &Spider{
	Name:            "拉勾-岗位",
	Description:     "拉勾上的全部岗位【https://www.lagou.com】",
	EnableCookie:    true,
	NotDefaultField: true,
	RuleTree:        lagouRuleTree,
}

var header = http.Header{}
var lagouRuleTree = &RuleTree{
	Root: func(ctx *Context) {
		ctx.AddQueue(&request.Request{
			Url:      positionURL,
			TryTimes: 10,
			Rule:     "requestList",
			Header:   header,
		})
	},
	Trunk: map[string]*Rule{
		"requestList": {
			ParseFunc: func(ctx *Context) {
				header.Set("Referer", ctx.Request.Url)
				nextSelection := ctx.GetDom().Find("div.pager_container").Find("a").Last()
				url, _ := nextSelection.Attr("href")
				if len(url) != 0 && strings.HasPrefix(url, "http") {
					ctx.AddQueue(&request.Request{
						Url:      url,
						TryTimes: 10,
						Rule:     "requestList",
						Priority: 1,
						Header:   header,
					})
				}
				ctx.Parse("outputResult")
			},
		},
		"outputResult": {
			ItemFields: []string{
				"岗位",
				"薪水",
				"工作地点",
				"公司",
			},
			ParseFunc: func(ctx *Context) {
				dom := ctx.GetDom()
				dom.Find("div.list_item_top").Each(func(i int, selection *goquery.Selection) {
					jobName := selection.Find("div.p_top").Find("h3").Text()
					city := selection.Find("div.p_top").Find("em").Text()
					city = strings.Split(city, "·")[0]
					salay := selection.Find("div.p_bot").Find("span.money").Text()
					company := selection.Find("div.company").Find("a").Text()
					ctx.Output(map[int]interface{}{
						0: jobName,
						1: salay,
						2: city,
						3: company,
					})
				})
			},
		},
	},
}
