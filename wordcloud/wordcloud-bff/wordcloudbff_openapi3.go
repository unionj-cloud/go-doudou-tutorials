package service

import "github.com/unionj-cloud/go-doudou/framework/http/onlinedoc"

func init() {
	onlinedoc.Oas = `{"openapi":"3.0.2","info":{"title":"WordcloudBff","version":"v20220301"},"servers":[{"url":"http://localhost:6060"}],"paths":{"/upload":{"post":{"requestBody":{"content":{"multipart/form-data":{"schema":{"$ref":"#/components/schemas/UploadReq"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/UploadResp"}}}}}}}},"components":{"schemas":{"UploadReq":{"title":"UploadReq","type":"object","properties":{"file":{"type":"string","format":"binary","description":"文本文件\ntext file"},"lang":{"type":"string","description":"语种\nlanguage"},"top":{"type":"integer","format":"int32","description":"取词频前top的词\nonly take top frequent words into word cloud"}},"required":["file","lang"]},"UploadResp":{"title":"UploadResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/UploadResult"}},"required":["data"]},"UploadResult":{"title":"UploadResult","type":"object","properties":{"imgUrl":{"type":"string","description":"词云图链接"},"srcUrl":{"type":"string","description":"原文件链接"},"taskId":{"type":"integer","format":"int32","description":"任务ID"}},"description":"\nUploadResult 上传文件接口返回对象","required":["taskId","srcUrl","imgUrl"]}}}}`
}
