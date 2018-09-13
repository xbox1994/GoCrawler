package config

const (
	// Elastic
	ElasticIndex = "dating_profile"
	ElasticType  = "zhenai"

	// Rpc Method
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"

	// Rate limiting
	Qps = 20
)
