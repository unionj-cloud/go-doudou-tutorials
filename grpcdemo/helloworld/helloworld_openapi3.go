package service

import "github.com/unionj-cloud/go-doudou/framework/http/onlinedoc"

func init() {
	onlinedoc.Oas = `{"openapi":"3.0.2","info":{"title":"Helloworld","description":"Helloworld 单体服务","version":"v20221001"},"servers":[{"url":"http://localhost:6060"}],"paths":{"/bye":{"post":{"description":"Bye 再见接口","requestBody":{"content":{"application/x-www-form-urlencoded":{"schema":{"$ref":"#/components/schemas/ByeReq"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/ByeResp"}}}}}}},"/greeting":{"post":{"description":"Greeting 问候接口","requestBody":{"content":{"application/x-www-form-urlencoded":{"schema":{"$ref":"#/components/schemas/GreetingReq"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GreetingResp"}}}}}}}},"components":{"schemas":{"ByeReq":{"title":"ByeReq","type":"object","properties":{"greeting":{"type":"string","description":"入参"}},"required":["greeting"]},"ByeResp":{"title":"ByeResp","type":"object","properties":{"data":{"type":"string","description":"出参"}},"required":["data"]},"GreetingReq":{"title":"GreetingReq","type":"object","properties":{"greeting":{"type":"string","description":"入参"}},"required":["greeting"]},"GreetingResp":{"title":"GreetingResp","type":"object","properties":{"data":{"type":"string","description":"出参"}},"required":["data"]},"Keyboard":{"title":"Keyboard","type":"object","properties":{"backlit":{"type":"boolean"},"layout":{"type":"string","default":"UNKNOWN","enum":["UNKNOWN","QWERTZ","AZERTY","QWERTY"]}},"required":["layout","backlit"]},"Order":{"title":"Order","type":"object","properties":{"col":{"type":"string"},"sort":{"type":"string"}},"required":["col","sort"]},"Page":{"title":"Page","type":"object","properties":{"orders":{"type":"array","items":{"$ref":"#/components/schemas/Order"},"description":"排序规则"},"pageNo":{"type":"integer","format":"int32","description":"页码"},"size":{"type":"integer","format":"int32","description":"每页行数"}},"required":["orders","pageNo","size"]},"PageFilter":{"title":"PageFilter","type":"object","properties":{"dept":{"type":"integer","format":"int32","description":"所属部门ID"},"name":{"type":"string","description":"真实姓名，前缀匹配"}},"required":["name","dept"]},"PageQuery":{"title":"PageQuery","type":"object","properties":{"filter":{"$ref":"#/components/schemas/PageFilter"},"page":{"$ref":"#/components/schemas/Page"}},"description":"分页筛选条件","required":["filter","page"]},"PageRet":{"title":"PageRet","type":"object","properties":{"hasNext":{"type":"boolean"},"items":{"type":"object"},"pageNo":{"type":"integer","format":"int32"},"pageSize":{"type":"integer","format":"int32"},"total":{"type":"integer","format":"int32"}},"required":["items","pageNo","pageSize","total","hasNext"]},"UserVo":{"title":"UserVo","type":"object","properties":{"dept":{"type":"string"},"id":{"type":"integer","format":"int32"},"name":{"type":"string"},"phone":{"type":"string"}},"required":["id","name","phone","dept"]}}}}`
}