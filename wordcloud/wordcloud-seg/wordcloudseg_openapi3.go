package service

import "github.com/unionj-cloud/go-doudou/framework/http/onlinedoc"

func init() {
	onlinedoc.Oas = `{"openapi":"3.0.2","info":{"title":"WordcloudSeg","version":"v20220307"},"servers":[{"url":"http://localhost:6060"}],"paths":{"/seg":{"post":{"requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/SegPayload"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/SegResp"}}}}}}}},"components":{"schemas":{"SegPayload":{"title":"SegPayload","type":"object","properties":{"lang":{"type":"string","description":"文本语言\n仅支持中文和英文\ntext language\nonly support zh and en"},"text":{"type":"string","description":"文本"}},"description":"\n","required":["text","lang"]},"SegResp":{"title":"SegResp","type":"object","properties":{"rs":{"$ref":"#/components/schemas/SegResult"}},"required":["rs"]},"SegResult":{"title":"SegResult","type":"object","properties":{"WordFreq":{"type":"array","items":{"type":"array","items":{"type":"object"}}}},"description":"\n","required":["WordFreq"]}}}}`
}
