package service

import "github.com/unionj-cloud/go-doudou/v2/framework/rest"

func init() {
	rest.Oas = `{"openapi":"3.0.2","info":{"title":"GoStats","description":"GoStats is a demo gRPC service developed by go-doudou","version":"v20221123"},"servers":[{"url":"http://localhost:6060"}],"paths":{"/largest/remainder":{"post":{"description":"LargestRemainder implements Largest Remainder Method https://en.wikipedia.org/wiki/Largest_remainder_method","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/PercentageReqVo"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/LargestRemainderResp"}}}}}}}},"components":{"schemas":{"LargestRemainderResp":{"title":"LargestRemainderResp","type":"object","properties":{"data":{"type":"array","items":{"$ref":"#/components/schemas/PercentageRespVo"}}},"required":["data"]},"PercentageReqVo":{"title":"PercentageReqVo","type":"object","properties":{"data":{"type":"array","items":{"$ref":"#/components/schemas/PercentageVo"},"description":"key value pairs"},"places":{"type":"integer","format":"int32","description":"digit number after dot"}},"description":"\nrequest vo","required":["data","places"]},"PercentageRespVo":{"title":"PercentageRespVo","type":"object","properties":{"key":{"type":"string"},"percent":{"type":"number","format":"double"},"percentFormatted":{"type":"string","description":"formatted percentage"},"value":{"type":"integer","format":"int32"}},"description":"result vo","required":["value","key","percent","percentFormatted"]},"PercentageVo":{"title":"PercentageVo","type":"object","properties":{"key":{"type":"string","description":"unique key for distinguishing something"},"value":{"type":"integer","format":"int32","description":"number for something"}},"description":"key value pair","required":["value","key"]}}}}`
}
