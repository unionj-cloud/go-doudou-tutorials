package service

import "github.com/unionj-cloud/go-doudou/framework/http/onlinedoc"

func init() {
	onlinedoc.Oas = `{"openapi":"3.0.2","info":{"title":"WordcloudSeg","version":"v20220223"},"servers":[{"url":"http://localhost:6060"}],"paths":{"/seg":{"post":{"requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/SegPayload"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/SegResp"}}}}}}}},"components":{"schemas":{"SegPayload":{"title":"SegPayload","type":"object","properties":{"text":{"type":"string"}},"description":"\n","required":["text"]},"SegResp":{"title":"SegResp","type":"object","properties":{"rs":{"$ref":"#/components/schemas/SegResult"}},"required":["rs"]},"SegResult":{"title":"SegResult","type":"object","properties":{"WordFreq":{"type":"object"}},"description":"\n","required":["WordFreq"]}}}}`
}
