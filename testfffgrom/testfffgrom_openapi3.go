package service

import "github.com/unionj-cloud/go-doudou/v2/framework/rest"

func init() {
	rest.Oas = `{"openapi":"3.0.2","info":{"title":"Testfffgrom","version":"v20230521"},"servers":[{"url":"http://localhost:6060"}],"paths":{"/t/author":{"post":{"description":"PostTAuthor mapped from table \u003ct_author\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/TAuthor"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PostTAuthorResp"}}}}}},"put":{"description":"PutTAuthor mapped from table \u003ct_author\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/TAuthor"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PutTAuthorResp"}}}}}}},"/t/author/{id}":{"get":{"description":"GetTAuthor_Id mapped from table \u003ct_author\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTAuthor_IdResp"}}}}}},"delete":{"description":"DeleteTAuthor_Id mapped from table \u003ct_author\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/DeleteTAuthor_IdResp"}}}}}}},"/t/authors":{"get":{"description":"GetTAuthors mapped from table \u003ct_author\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"parameter","in":"query","required":true,"schema":{"$ref":"#/components/schemas/Parameter"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTAuthorsResp"}}}}}}},"/t/client":{"post":{"description":"PostTClient mapped from table \u003ct_client\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/TClient"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PostTClientResp"}}}}}},"put":{"description":"PutTClient mapped from table \u003ct_client\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/TClient"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PutTClientResp"}}}}}}},"/t/client/{id}":{"get":{"description":"GetTClient_Id mapped from table \u003ct_client\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTClient_IdResp"}}}}}},"delete":{"description":"DeleteTClient_Id mapped from table \u003ct_client\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/DeleteTClient_IdResp"}}}}}}},"/t/clients":{"get":{"description":"GetTClients mapped from table \u003ct_client\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"parameter","in":"query","required":true,"schema":{"$ref":"#/components/schemas/Parameter"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTClientsResp"}}}}}}},"/t/dept":{"post":{"description":"PostTDept mapped from table \u003ct_dept\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/TDept"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PostTDeptResp"}}}}}},"put":{"description":"PutTDept mapped from table \u003ct_dept\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/TDept"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PutTDeptResp"}}}}}}},"/t/dept/{id}":{"get":{"description":"GetTDept_Id mapped from table \u003ct_dept\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTDept_IdResp"}}}}}},"delete":{"description":"DeleteTDept_Id mapped from table \u003ct_dept\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/DeleteTDept_IdResp"}}}}}}},"/t/depts":{"get":{"description":"GetTDepts mapped from table \u003ct_dept\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"parameter","in":"query","required":true,"schema":{"$ref":"#/components/schemas/Parameter"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTDeptsResp"}}}}}}},"/t/invalid/token":{"post":{"description":"PostTInvalidToken mapped from table \u003ct_invalid_token\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/TInvalidToken"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PostTInvalidTokenResp"}}}}}},"put":{"description":"PutTInvalidToken mapped from table \u003ct_invalid_token\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/TInvalidToken"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PutTInvalidTokenResp"}}}}}}},"/t/invalid/token/{id}":{"get":{"description":"GetTInvalidToken_Id mapped from table \u003ct_invalid_token\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTInvalidToken_IdResp"}}}}}},"delete":{"description":"DeleteTInvalidToken_Id mapped from table \u003ct_invalid_token\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/DeleteTInvalidToken_IdResp"}}}}}}},"/t/invalid/tokens":{"get":{"description":"GetTInvalidTokens mapped from table \u003ct_invalid_token\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"parameter","in":"query","required":true,"schema":{"$ref":"#/components/schemas/Parameter"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTInvalidTokensResp"}}}}}}},"/t/jjj":{"post":{"description":"PostTJjj mapped from table \u003ct_jjj\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/TJjj"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PostTJjjResp"}}}}}},"put":{"description":"PutTJjj mapped from table \u003ct_jjj\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/TJjj"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PutTJjjResp"}}}}}}},"/t/jjj/{id}":{"get":{"description":"GetTJjj_Id mapped from table \u003ct_jjj\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTJjj_IdResp"}}}}}},"delete":{"description":"DeleteTJjj_Id mapped from table \u003ct_jjj\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/DeleteTJjj_IdResp"}}}}}}},"/t/jjjs":{"get":{"description":"GetTJjjs mapped from table \u003ct_jjj\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"parameter","in":"query","required":true,"schema":{"$ref":"#/components/schemas/Parameter"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTJjjsResp"}}}}}}},"/t/post":{"post":{"description":"PostTPost mapped from table \u003ct_post\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/TPost"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PostTPostResp"}}}}}},"put":{"description":"PutTPost mapped from table \u003ct_post\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/TPost"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PutTPostResp"}}}}}}},"/t/post/{id}":{"get":{"description":"GetTPost_Id mapped from table \u003ct_post\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTPost_IdResp"}}}}}},"delete":{"description":"DeleteTPost_Id mapped from table \u003ct_post\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/DeleteTPost_IdResp"}}}}}}},"/t/posts":{"get":{"description":"GetTPosts mapped from table \u003ct_post\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"parameter","in":"query","required":true,"schema":{"$ref":"#/components/schemas/Parameter"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTPostsResp"}}}}}}},"/t/user":{"post":{"description":"PostTUser mapped from table \u003ct_user\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/TUser"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PostTUserResp"}}}}}},"put":{"description":"PutTUser mapped from table \u003ct_user\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/TUser"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PutTUserResp"}}}}}}},"/t/user/{id}":{"get":{"description":"GetTUser_Id mapped from table \u003ct_user\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTUser_IdResp"}}}}}},"delete":{"description":"DeleteTUser_Id mapped from table \u003ct_user\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/DeleteTUser_IdResp"}}}}}}},"/t/users":{"get":{"description":"GetTUsers mapped from table \u003ct_user\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"parameter","in":"query","required":true,"schema":{"$ref":"#/components/schemas/Parameter"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTUsersResp"}}}}}}},"/t/word/cloud/task":{"post":{"description":"PostTWordCloudTask mapped from table \u003ct_word_cloud_task\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/TWordCloudTask"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PostTWordCloudTaskResp"}}}}}},"put":{"description":"PutTWordCloudTask mapped from table \u003ct_word_cloud_task\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/TWordCloudTask"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PutTWordCloudTaskResp"}}}}}}},"/t/word/cloud/task/{id}":{"get":{"description":"GetTWordCloudTask_Id mapped from table \u003ct_word_cloud_task\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTWordCloudTask_IdResp"}}}}}},"delete":{"description":"DeleteTWordCloudTask_Id mapped from table \u003ct_word_cloud_task\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"integer","format":"int32"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/DeleteTWordCloudTask_IdResp"}}}}}}},"/t/word/cloud/tasks":{"get":{"description":"GetTWordCloudTasks mapped from table \u003ct_word_cloud_task\u003e\nCode generated by gorm.io/gen for go-doudou. DO NOT EDIT.","parameters":[{"name":"parameter","in":"query","required":true,"schema":{"$ref":"#/components/schemas/Parameter"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetTWordCloudTasksResp"}}}}}}}},"components":{"schemas":{"DeleteTAuthor_IdResp":{"title":"DeleteTAuthor_IdResp","type":"object"},"DeleteTClient_IdResp":{"title":"DeleteTClient_IdResp","type":"object"},"DeleteTDept_IdResp":{"title":"DeleteTDept_IdResp","type":"object"},"DeleteTInvalidToken_IdResp":{"title":"DeleteTInvalidToken_IdResp","type":"object"},"DeleteTJjj_IdResp":{"title":"DeleteTJjj_IdResp","type":"object"},"DeleteTPost_IdResp":{"title":"DeleteTPost_IdResp","type":"object"},"DeleteTUser_IdResp":{"title":"DeleteTUser_IdResp","type":"object"},"DeleteTWordCloudTask_IdResp":{"title":"DeleteTWordCloudTask_IdResp","type":"object"},"GetTAuthor_IdResp":{"title":"GetTAuthor_IdResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/TAuthor"}},"required":["data"]},"GetTAuthorsResp":{"title":"GetTAuthorsResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/Page"}},"required":["data"]},"GetTClient_IdResp":{"title":"GetTClient_IdResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/TClient"}},"required":["data"]},"GetTClientsResp":{"title":"GetTClientsResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/Page"}},"required":["data"]},"GetTDept_IdResp":{"title":"GetTDept_IdResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/TDept"}},"required":["data"]},"GetTDeptsResp":{"title":"GetTDeptsResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/Page"}},"required":["data"]},"GetTInvalidToken_IdResp":{"title":"GetTInvalidToken_IdResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/TInvalidToken"}},"required":["data"]},"GetTInvalidTokensResp":{"title":"GetTInvalidTokensResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/Page"}},"required":["data"]},"GetTJjj_IdResp":{"title":"GetTJjj_IdResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/TJjj"}},"required":["data"]},"GetTJjjsResp":{"title":"GetTJjjsResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/Page"}},"required":["data"]},"GetTPost_IdResp":{"title":"GetTPost_IdResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/TPost"}},"required":["data"]},"GetTPostsResp":{"title":"GetTPostsResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/Page"}},"required":["data"]},"GetTUser_IdResp":{"title":"GetTUser_IdResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/TUser"}},"required":["data"]},"GetTUsersResp":{"title":"GetTUsersResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/Page"}},"required":["data"]},"GetTWordCloudTask_IdResp":{"title":"GetTWordCloudTask_IdResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/TWordCloudTask"}},"required":["data"]},"GetTWordCloudTasksResp":{"title":"GetTWordCloudTasksResp","type":"object","properties":{"data":{"$ref":"#/components/schemas/Page"}},"required":["data"]},"Page":{"title":"Page","type":"object","properties":{"first":{"type":"boolean"},"items":{"type":"array","items":{"type":"object"}},"last":{"type":"boolean"},"maxPage":{"type":"integer","format":"int64"},"page":{"type":"integer","format":"int64"},"size":{"type":"integer","format":"int64"},"total":{"type":"integer","format":"int64"},"totalPages":{"type":"integer","format":"int64"},"visible":{"type":"integer","format":"int64"}},"description":"\n\n\n\n\n\n\n\nPage result wrapper","required":["items","page","size","maxPage","totalPages","total","last","first","visible"]},"Parameter":{"title":"Parameter","type":"object","properties":{"fields":{"type":"string"},"filters":{"type":"array","items":{"type":"object"}},"order":{"type":"string"},"page":{"type":"string"},"size":{"type":"string"},"sort":{"type":"string"}},"description":"\n\n\n\n\n\n\n\nParameter struct","required":["page","size","sort","order","fields","filters"]},"PostTAuthorResp":{"title":"PostTAuthorResp","type":"object","properties":{"data":{"type":"integer","format":"int32"}},"required":["data"]},"PostTClientResp":{"title":"PostTClientResp","type":"object","properties":{"data":{"type":"integer","format":"int32"}},"required":["data"]},"PostTDeptResp":{"title":"PostTDeptResp","type":"object","properties":{"data":{"type":"integer","format":"int32"}},"required":["data"]},"PostTInvalidTokenResp":{"title":"PostTInvalidTokenResp","type":"object","properties":{"data":{"type":"integer","format":"int32"}},"required":["data"]},"PostTJjjResp":{"title":"PostTJjjResp","type":"object","properties":{"data":{"type":"integer","format":"int32"}},"required":["data"]},"PostTPostResp":{"title":"PostTPostResp","type":"object","properties":{"data":{"type":"integer","format":"int32"}},"required":["data"]},"PostTUserResp":{"title":"PostTUserResp","type":"object","properties":{"data":{"type":"integer","format":"int32"}},"required":["data"]},"PostTWordCloudTaskResp":{"title":"PostTWordCloudTaskResp","type":"object","properties":{"data":{"type":"integer","format":"int32"}},"required":["data"]},"PutTAuthorResp":{"title":"PutTAuthorResp","type":"object"},"PutTClientResp":{"title":"PutTClientResp","type":"object"},"PutTDeptResp":{"title":"PutTDeptResp","type":"object"},"PutTInvalidTokenResp":{"title":"PutTInvalidTokenResp","type":"object"},"PutTJjjResp":{"title":"PutTJjjResp","type":"object"},"PutTPostResp":{"title":"PutTPostResp","type":"object"},"PutTUserResp":{"title":"PutTUserResp","type":"object"},"PutTWordCloudTaskResp":{"title":"PutTWordCloudTaskResp","type":"object"},"TAuthor":{"title":"TAuthor","type":"object","properties":{"book":{"type":"string"},"deleteAt":{"type":"object"},"id":{"type":"integer","format":"int32"}},"description":"\n\n\nTAuthor mapped from table \u003ct_author\u003e","required":["id","deleteAt"]},"TClient":{"title":"TClient","type":"object","properties":{"createAt":{"type":"string","format":"date-time"},"deleteAt":{"type":"object"},"encryptedSecret":{"type":"string"},"id":{"type":"integer","format":"int32"},"name":{"type":"string"},"publicKey":{"type":"string"},"updateAt":{"type":"string","format":"date-time"}},"description":"\n\n\nTClient mapped from table \u003ct_client\u003e","required":["id","name","publicKey","encryptedSecret","deleteAt"]},"TDept":{"title":"TDept","type":"object","properties":{"deletedAt":{"type":"object"},"id":{"type":"integer","format":"int32"},"name":{"type":"string"}},"description":"\n\n\nTDept mapped from table \u003ct_dept\u003e","required":["id","name","deletedAt"]},"TInvalidToken":{"title":"TInvalidToken","type":"object","properties":{"createAt":{"type":"string","format":"date-time"},"deleteAt":{"type":"object"},"expireAt":{"type":"string","format":"date-time"},"id":{"type":"integer","format":"int32"},"token":{"type":"string"},"updateAt":{"type":"string","format":"date-time"}},"description":"\n\n\nTInvalidToken mapped from table \u003ct_invalid_token\u003e","required":["id","token","deleteAt"]},"TJjj":{"title":"TJjj","type":"object","properties":{"deletedAt":{"type":"object"},"id":{"type":"integer","format":"int32"},"jobName":{"type":"string"}},"description":"\n\n\nTJjj mapped from table \u003ct_jjj\u003e","required":["id","deletedAt"]},"TPost":{"title":"TPost","type":"object","properties":{"deleteAt":{"type":"object"},"id":{"type":"integer","format":"int32"},"title":{"type":"string"}},"description":"\n\n\nTPost mapped from table \u003ct_post\u003e","required":["id","deleteAt"]},"TUser":{"title":"TUser","type":"object","properties":{"avatar":{"type":"string"},"createAt":{"type":"string","format":"date-time"},"deleteAt":{"type":"object"},"dept":{"type":"string"},"id":{"type":"integer","format":"int32"},"name":{"type":"string"},"password":{"type":"string"},"phone":{"type":"string"},"updateAt":{"type":"string","format":"date-time"},"username":{"type":"string"}},"description":"\n\n\nTUser mapped from table \u003ct_user\u003e","required":["id","username","password","name","phone","dept","avatar","deleteAt"]},"TWordCloudTask":{"title":"TWordCloudTask","type":"object","properties":{"createAt":{"type":"string","format":"date-time"},"deleteAt":{"type":"object"},"error":{"type":"string"},"id":{"type":"integer","format":"int32"},"imgURL":{"type":"string"},"lang":{"type":"string"},"srcURL":{"type":"string"},"status":{"type":"integer","format":"int32"},"top":{"type":"integer","format":"int32"},"updateAt":{"type":"string","format":"date-time"},"userID":{"type":"integer","format":"int32"}},"description":"\n\n\nTWordCloudTask mapped from table \u003ct_word_cloud_task\u003e","required":["id","srcURL","imgURL","lang","top","status","error","userID","deleteAt"]},"User":{"title":"User","type":"object","properties":{"dept":{"type":"string"},"id":{"type":"integer","format":"int32"},"name":{"type":"string"},"phone":{"type":"string"}},"required":["id","name","phone","dept"]}}}}`
}
