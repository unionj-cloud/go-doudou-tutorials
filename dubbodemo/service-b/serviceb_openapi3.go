package service

import "github.com/unionj-cloud/go-doudou/v2/framework/rest"

func init() {
	rest.Oas = `{"openapi":"3.0.2","info":{"title":"ServiceB","version":"v20230420"},"servers":[{"url":"http://localhost:6060"}],"paths":{"/dept/by/id":{"get":{"parameters":[{"name":"deptId","in":"query","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetDeptByIdResp"}}}}}}}},"components":{"schemas":{"DeptDto":{"title":"DeptDto","type":"object","properties":{"id":{"type":"integer","format":"int32"},"name":{"type":"string"},"staffTotal":{"type":"integer","format":"int32"}},"description":"\n","required":["id","name","staffTotal"]},"GetDeptByIdResp":{"title":"GetDeptByIdResp","type":"object","properties":{"dept":{"$ref":"#/components/schemas/DeptDto"}},"required":["dept"]},"Order":{"title":"Order","type":"object","properties":{"col":{"type":"string"},"sort":{"type":"string"}},"required":["col","sort"]},"Page":{"title":"Page","type":"object","properties":{"orders":{"type":"array","items":{"$ref":"#/components/schemas/Order"},"description":"排序规则"},"pageNo":{"type":"integer","format":"int32","description":"页码"},"size":{"type":"integer","format":"int32","description":"每页行数"}},"required":["orders","pageNo","size"]},"PageFilter":{"title":"PageFilter","type":"object","properties":{"dept":{"type":"integer","format":"int32","description":"所属部门ID"},"name":{"type":"string","description":"真实姓名，前缀匹配"}},"required":["name","dept"]},"PageQuery":{"title":"PageQuery","type":"object","properties":{"filter":{"$ref":"#/components/schemas/PageFilter"},"page":{"$ref":"#/components/schemas/Page"}},"description":"分页筛选条件","required":["filter","page"]},"PageRet":{"title":"PageRet","type":"object","properties":{"hasNext":{"type":"boolean"},"items":{"type":"object"},"pageNo":{"type":"integer","format":"int32"},"pageSize":{"type":"integer","format":"int32"},"total":{"type":"integer","format":"int32"}},"required":["items","pageNo","pageSize","total","hasNext"]},"UserDto":{"title":"UserDto","type":"object","properties":{"dept":{"type":"string"},"id":{"type":"integer","format":"int32"},"name":{"type":"string"},"phone":{"type":"string"}},"required":["id","name","phone","dept"]}}}}`
}