// Author: xufei
// Date: 2018/12/18

package config

const (
	ElasticSearchDB          = "dating_profile"
	ElasticSearchZhenAiTable = "zhenai"

	// RPC 端口
	PersistRpcPort = 1235
	WorkerPort0    = 9000

	// 解析器名称
	ProfileParser = "ProfileParser"
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"

	// RPC服务名称
	CrawlServiceRpc = "CrawlService.Process"
)
