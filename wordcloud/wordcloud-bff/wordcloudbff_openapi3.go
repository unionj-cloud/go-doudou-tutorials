package service

import "github.com/unionj-cloud/go-doudou/framework/http/onlinedoc"

func init() {
	onlinedoc.Oas = `{"openapi":"3.0.2","info":{"title":"WordcloudBff","version":"v20220315"},"servers":[{"url":"http://localhost:6060"}],"paths":{"/task/page":{"get":{"description":"词云图任务分页接口\nword cloud task pagination","parameters":[{"name":"page","in":"query","required":true,"schema":{"type":"integer","format":"int32"}},{"name":"pageSize","in":"query","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTaskPageResp"}}}}}}},"/upload":{"post":{"description":"上传文本文件生成词云图接口\nupload text file to generate word cloud image","requestBody":{"content":{"multipart/form-data":{"schema":{"$ref":"#/components/schemas/UploadReq"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/UploadResp"}}}}}}}},"components":{"schemas":{"GetTaskPageResp":{"title":"GetTaskPageResp","type":"object","properties":{"result":{"$ref":"#/components/schemas/TaskPageRet"}},"required":["result"]},"Order":{"title":"Order","type":"object","properties":{"col":{"type":"string"},"sort":{"type":"string"}},"required":["col","sort"]},"Page":{"title":"Page","type":"object","properties":{"orders":{"type":"array","items":{"$ref":"#/components/schemas/Order"},"description":"排序规则"},"pageNo":{"type":"integer","format":"int32","description":"页码"},"size":{"type":"integer","format":"int32","description":"每页行数"}},"required":["orders","pageNo","size"]},"PageQuery":{"title":"PageQuery","type":"object","properties":{"page":{"$ref":"#/components/schemas/Page"}},"description":"分页筛选条件","required":["page"]},"PageRet":{"title":"PageRet","type":"object","properties":{"hasNext":{"type":"boolean"},"items":{"type":"object"},"pageNo":{"type":"integer","format":"int32"},"pageSize":{"type":"integer","format":"int32"},"total":{"type":"integer","format":"int32"}},"required":["items","pageNo","pageSize","total","hasNext"]},"TaskPageRet":{"title":"TaskPageRet","type":"object","properties":{"hasNext":{"type":"boolean"},"items":{"type":"array","items":{"$ref":"#/components/schemas/TaskVo"}},"pageNo":{"type":"integer","format":"int32"},"pageSize":{"type":"integer","format":"int32"},"total":{"type":"integer","format":"int32"}},"description":"分页结果\npagination result\n","required":["items","hasNext","pageNo","pageSize","total"]},"TaskVo":{"title":"TaskVo","type":"object","properties":{"createAt":{"type":"string"},"error":{"type":"string"},"id":{"type":"integer","format":"int32"},"imgUrl":{"type":"string"},"lang":{"type":"string"},"srcUrl":{"type":"string"},"status":{"type":"integer","format":"int32"},"top":{"type":"integer","format":"int32"},"username":{"type":"string"}},"required":["id","srcUrl","imgUrl","lang","top","status","error","username","createAt"]},"UploadReq":{"title":"UploadReq","type":"object","properties":{"file":{"type":"string","format":"binary","description":"文本文件\ntext file"},"lang":{"type":"string","description":"语种\nlanguage"},"top":{"type":"integer","format":"int32","description":"取词频前top的词\nonly take top frequent words into word cloud"}},"required":["file","lang"]},"UploadResp":{"title":"UploadResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/UploadResult"}},"required":["data"]},"UploadResult":{"title":"UploadResult","type":"object","properties":{"imgUrl":{"type":"string","description":"词云图链接"},"srcUrl":{"type":"string","description":"原文件链接"},"taskId":{"type":"integer","format":"int32","description":"任务ID"}},"description":"\nUploadResult 上传文件接口返回对象","required":["taskId","srcUrl","imgUrl"]}}}}`
}
