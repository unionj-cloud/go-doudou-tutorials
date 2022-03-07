package service

import "github.com/unionj-cloud/go-doudou/framework/http/onlinedoc"

func init() {
	onlinedoc.Oas = `{"openapi":"3.0.2","info":{"title":"WordcloudMaker","version":"v20220307"},"servers":[{"url":"http://localhost:6060"}],"paths":{"/make":{"post":{"description":"Make 生成词云图接口\nmake word cloud image","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/MakePayload"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/MakeResp"}}}}}}}},"components":{"schemas":{"MakePayload":{"title":"MakePayload","type":"object","properties":{"lang":{"type":"string","description":"文本语言\ntext language"},"srcUrl":{"type":"string","description":"文本文件公开下载链接\ntext file public download url"},"top":{"type":"integer","format":"int32","description":"只取词频前Top个词\nonly want Top frequent words"}},"description":"\n","required":["srcUrl","lang","top"]},"MakeResp":{"title":"MakeResp","type":"object","properties":{"data":{"type":"string","description":"词云图链接\nword cloud image url"}},"required":["data"]}}}}`
}
