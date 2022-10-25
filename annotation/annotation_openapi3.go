package service

import "github.com/unionj-cloud/go-doudou/v2/framework/rest"

func init() {
	rest.Oas = `{"openapi":"3.0.2","info":{"title":"Annotation","version":"v20221024"},"servers":[{"url":"http://localhost:6060"}],"paths":{"/admin":{"get":{"description":"此接口只有管理员有权访问\n@role(ADMIN)","responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetAdminResp"}}}}}}},"/guest":{"get":{"description":"此接口可公开访问，无需校验登录和权限","responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetGuestResp"}}}}}}},"/user":{"get":{"description":"此接口只有登录用户有权访问\n@role(USER,ADMIN)","responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetUserResp"}}}}}}}},"components":{"schemas":{"Auth":{"title":"Auth","type":"object","properties":{"Pass":{"type":"string"},"User":{"type":"string"}},"required":["User","Pass"]},"GetAdminResp":{"title":"GetAdminResp","type":"object","properties":{"data":{"type":"string"}},"required":["data"]},"GetGuestResp":{"title":"GetGuestResp","type":"object","properties":{"data":{"type":"string"}},"required":["data"]},"GetUserResp":{"title":"GetUserResp","type":"object","properties":{"data":{"type":"string"}},"required":["data"]},"Order":{"title":"Order","type":"object","properties":{"Col":{"type":"string"},"Sort":{"type":"string"}},"required":["Col","Sort"]},"Page":{"title":"Page","type":"object","properties":{"Orders":{"type":"array","items":{"$ref":"#/components/schemas/Order"},"description":"排序规则"},"PageNo":{"type":"integer","format":"int32","description":"页码"},"Size":{"type":"integer","format":"int32","description":"每页行数"}},"required":["Orders","PageNo","Size"]},"PageFilter":{"title":"PageFilter","type":"object","properties":{"Dept":{"type":"integer","format":"int32","description":"所属部门ID"},"Name":{"type":"string","description":"真实姓名，前缀匹配"}},"required":["Name","Dept"]},"PageQuery":{"title":"PageQuery","type":"object","properties":{"Filter":{"$ref":"#/components/schemas/PageFilter"},"Page":{"$ref":"#/components/schemas/Page"}},"description":"分页筛选条件","required":["Filter","Page"]},"PageRet":{"title":"PageRet","type":"object","properties":{"HasNext":{"type":"boolean"},"Items":{"type":"object"},"PageNo":{"type":"integer","format":"int32"},"PageSize":{"type":"integer","format":"int32"},"Total":{"type":"integer","format":"int32"}},"required":["Items","PageNo","PageSize","Total","HasNext"]},"UserVo":{"title":"UserVo","type":"object","properties":{"Dept":{"type":"string"},"Id":{"type":"integer","format":"int32"},"Name":{"type":"string"},"Phone":{"type":"string"}},"required":["Id","Name","Phone","Dept"]}}}}`
}
