package service

import "github.com/unionj-cloud/go-doudou/v2/framework/rest"

func init() {
	rest.Oas = `{"openapi":"3.0.2","info":{"title":"Bistream","version":"v20221223"},"servers":[{"url":"http://localhost:6060"}],"paths":{"/survey":{"post":{"requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/Service"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/SurveyResp"}}}}}}}},"components":{"schemas":{"Service":{"title":"Service","type":"object","properties":{"name":{"type":"string"},"restAddr":{"type":"string"},"rpcAddr":{"type":"string"},"weight":{"type":"integer","format":"int32"}},"description":"\n\n","required":["name","rpcAddr","restAddr","weight"]},"SurveyResp":{"title":"SurveyResp","type":"object","properties":{"stream1":{"$ref":"#/components/schemas/Service"}},"required":["stream1"]}}}}`
}
