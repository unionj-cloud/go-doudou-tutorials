package service

import "github.com/unionj-cloud/go-doudou/v2/framework/rest"

func init() {
	rest.Oas = `{"openapi":"3.0.2","info":{"title":"Gin2Doudou","description":"Gin2Doudou gin-admin\nRBAC scaffolding based on GIN + GORM + CASBIN + WIRE.\n\n8.1.0\ntiannianshou@gmail.com\n\n","version":"v20230129"},"servers":[{"url":"http://localhost:6060"}],"paths":{"/api/v1/menus":{"get":{"description":"GetApiV1Menus 查询数据","parameters":[{"name":"current","in":"query","description":"分页索引\nrequired","required":true,"schema":{"type":"integer","format":"int32","description":"分页索引\nrequired"}},{"name":"pageSize","in":"query","description":"分页大小\nrequired","required":true,"schema":{"type":"integer","format":"int32","description":"分页大小\nrequired"}},{"name":"queryValue","in":"query","description":"查询值","schema":{"type":"string","description":"查询值"}},{"name":"status","in":"query","description":"状态(1:启用 2:禁用)","schema":{"type":"integer","format":"int32","description":"状态(1:启用 2:禁用)"}},{"name":"isShow","in":"query","description":"是否显示(1:显示 2:隐藏)","schema":{"type":"integer","format":"int32","description":"是否显示(1:显示 2:隐藏)"}},{"name":"parentID","in":"query","description":"父级ID","schema":{"type":"integer","format":"int32","description":"父级ID"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetApiV1MenusResp"}}}}}},"post":{"description":"ApiV1Menus 创建数据","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/SchemaMenu"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/ApiV1MenusResp"}}}}}}},"/api/v1/menus/{id}":{"get":{"description":"GetApiV1Menus_Id 查询指定数据","parameters":[{"name":"id","in":"path","description":"唯一标识\nrequired","required":true,"schema":{"type":"integer","format":"int32","description":"唯一标识\nrequired"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetApiV1Menus_IdResp"}}}}}},"put":{"description":"PutApiV1Menus_Id 更新数据","parameters":[{"name":"id","in":"path","description":"唯一标识\nrequired","required":true,"schema":{"type":"integer","format":"int32","description":"唯一标识\nrequired"}}],"requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/SchemaMenu"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PutApiV1Menus_IdResp"}}}}}},"delete":{"description":"DeleteApiV1Menus_Id 删除数据","parameters":[{"name":"id","in":"path","description":"唯一标识\nrequired","required":true,"schema":{"type":"integer","format":"int32","description":"唯一标识\nrequired"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/DeleteApiV1Menus_IdResp"}}}}}}},"/api/v1/menustree":{"get":{"description":"GetApiV1Menustree 查询菜单树","parameters":[{"name":"status","in":"query","description":"状态(1:启用 2:禁用)","schema":{"type":"integer","format":"int32","description":"状态(1:启用 2:禁用)"}},{"name":"parentID","in":"query","description":"父级ID","schema":{"type":"integer","format":"int32","description":"父级ID"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetApiV1MenustreeResp"}}}}}}},"/api/v1/pub/current/menutree":{"get":{"description":"GetApiV1PubCurrentMenutree 查询当前用户菜单树","responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetApiV1PubCurrentMenutreeResp"}}}}}}},"/api/v1/pub/current/password":{"put":{"description":"PutApiV1PubCurrentPassword 更新个人密码","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/SchemaUpdatePasswordParam"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PutApiV1PubCurrentPasswordResp"}}}}}}},"/api/v1/pub/current/user":{"get":{"description":"GetApiV1PubCurrentUser 获取当前用户信息","responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetApiV1PubCurrentUserResp"}}}}}}},"/api/v1/pub/login":{"post":{"description":"ApiV1PubLogin 用户登录","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/SchemaLoginParam"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/ApiV1PubLoginResp"}}}}}}},"/api/v1/pub/login/captcha":{"get":{"description":"GetApiV1PubLoginCaptcha 响应图形验证码","parameters":[{"name":"id","in":"query","description":"验证码ID\nrequired","required":true,"schema":{"type":"string","description":"验证码ID\nrequired"}},{"name":"reload","in":"query","description":"重新加载","schema":{"type":"string","description":"重新加载"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetApiV1PubLoginCaptchaResp"}}}}}}},"/api/v1/pub/login/captchaid":{"get":{"description":"GetApiV1PubLoginCaptchaid 获取验证码信息","responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetApiV1PubLoginCaptchaidResp"}}}}}}},"/api/v1/pub/login/exit":{"post":{"description":"ApiV1PubLoginExit 用户登出","requestBody":{"content":{"application/x-www-form-urlencoded":{"schema":{"$ref":"#/components/schemas/ApiV1PubLoginExitReq"}}}},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/ApiV1PubLoginExitResp"}}}}}}},"/api/v1/pub/refreshtoken":{"post":{"description":"ApiV1PubRefreshtoken 刷新令牌","requestBody":{"content":{"application/x-www-form-urlencoded":{"schema":{"$ref":"#/components/schemas/ApiV1PubRefreshtokenReq"}}}},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/ApiV1PubRefreshtokenResp"}}}}}}},"/api/v1/roles":{"get":{"description":"GetApiV1Roles 查询数据","parameters":[{"name":"current","in":"query","description":"分页索引\nrequired","required":true,"schema":{"type":"integer","format":"int32","description":"分页索引\nrequired"}},{"name":"pageSize","in":"query","description":"分页大小\nrequired","required":true,"schema":{"type":"integer","format":"int32","description":"分页大小\nrequired"}},{"name":"queryValue","in":"query","description":"查询值","schema":{"type":"string","description":"查询值"}},{"name":"status","in":"query","description":"状态(1:启用 2:禁用)","schema":{"type":"integer","format":"int32","description":"状态(1:启用 2:禁用)"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetApiV1RolesResp"}}}}}},"post":{"description":"ApiV1Roles 创建数据","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/SchemaRole"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/ApiV1RolesResp"}}}}}}},"/api/v1/roles/{id}":{"get":{"description":"GetApiV1Roles_Id 查询指定数据","parameters":[{"name":"id","in":"path","description":"唯一标识\nrequired","required":true,"schema":{"type":"integer","format":"int32","description":"唯一标识\nrequired"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetApiV1Roles_IdResp"}}}}}},"put":{"description":"PutApiV1Roles_Id 更新数据","parameters":[{"name":"id","in":"path","description":"唯一标识\nrequired","required":true,"schema":{"type":"integer","format":"int32","description":"唯一标识\nrequired"}}],"requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/SchemaRole"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PutApiV1Roles_IdResp"}}}}}},"delete":{"description":"DeleteApiV1Roles_Id 删除数据","parameters":[{"name":"id","in":"path","description":"唯一标识\nrequired","required":true,"schema":{"type":"integer","format":"int32","description":"唯一标识\nrequired"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/DeleteApiV1Roles_IdResp"}}}}}}},"/api/v1/rolesselect":{"get":{"description":"GetApiV1Rolesselect 查询选择数据","parameters":[{"name":"queryValue","in":"query","description":"查询值","schema":{"type":"string","description":"查询值"}},{"name":"status","in":"query","description":"状态(1:启用 2:禁用)","schema":{"type":"integer","format":"int32","description":"状态(1:启用 2:禁用)"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetApiV1RolesselectResp"}}}}}}},"/api/v1/users":{"get":{"description":"GetApiV1Users 查询数据","parameters":[{"name":"current","in":"query","description":"分页索引\nrequired","required":true,"schema":{"type":"integer","format":"int32","description":"分页索引\nrequired"}},{"name":"pageSize","in":"query","description":"分页大小\nrequired","required":true,"schema":{"type":"integer","format":"int32","description":"分页大小\nrequired"}},{"name":"queryValue","in":"query","description":"查询值","schema":{"type":"string","description":"查询值"}},{"name":"roleIDs","in":"query","description":"角色ID(多个以英文逗号分隔)","schema":{"type":"string","description":"角色ID(多个以英文逗号分隔)"}},{"name":"status","in":"query","description":"状态(1:启用 2:停用)","schema":{"type":"integer","format":"int32","description":"状态(1:启用 2:停用)"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetApiV1UsersResp"}}}}}},"post":{"description":"ApiV1Users 创建数据","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/SchemaUser"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/ApiV1UsersResp"}}}}}}},"/api/v1/users/{id}":{"get":{"description":"GetApiV1Users_Id 查询指定数据","parameters":[{"name":"id","in":"path","description":"唯一标识\nrequired","required":true,"schema":{"type":"integer","format":"int32","description":"唯一标识\nrequired"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetApiV1Users_IdResp"}}}}}},"put":{"description":"PutApiV1Users_Id 更新数据","parameters":[{"name":"id","in":"path","description":"唯一标识\nrequired","required":true,"schema":{"type":"integer","format":"int32","description":"唯一标识\nrequired"}}],"requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/SchemaUser"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PutApiV1Users_IdResp"}}}}}},"delete":{"description":"DeleteApiV1Users_Id 删除数据","parameters":[{"name":"id","in":"path","description":"唯一标识\nrequired","required":true,"schema":{"type":"integer","format":"int32","description":"唯一标识\nrequired"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/DeleteApiV1Users_IdResp"}}}}}}}},"components":{"schemas":{"ApiV1MenusResp":{"title":"ApiV1MenusResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaIDResult"}},"required":["ret"]},"ApiV1PubLoginExitReq":{"title":"ApiV1PubLoginExitReq","type":"object"},"ApiV1PubLoginExitResp":{"title":"ApiV1PubLoginExitResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaStatusResult"}},"required":["ret"]},"ApiV1PubLoginResp":{"title":"ApiV1PubLoginResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaLoginTokenInfo"}},"required":["ret"]},"ApiV1PubRefreshtokenReq":{"title":"ApiV1PubRefreshtokenReq","type":"object"},"ApiV1PubRefreshtokenResp":{"title":"ApiV1PubRefreshtokenResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaLoginTokenInfo"}},"required":["ret"]},"ApiV1RolesResp":{"title":"ApiV1RolesResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaIDResult"}},"required":["ret"]},"ApiV1UsersResp":{"title":"ApiV1UsersResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaIDResult"}},"required":["ret"]},"DeleteApiV1Menus_IdResp":{"title":"DeleteApiV1Menus_IdResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaStatusResult"}},"required":["ret"]},"DeleteApiV1Roles_IdResp":{"title":"DeleteApiV1Roles_IdResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaStatusResult"}},"required":["ret"]},"DeleteApiV1Users_IdResp":{"title":"DeleteApiV1Users_IdResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaStatusResult"}},"required":["ret"]},"GetApiV1MenusResp":{"title":"GetApiV1MenusResp","type":"object","properties":{"ret":{"type":"object"}},"required":["ret"]},"GetApiV1Menus_IdResp":{"title":"GetApiV1Menus_IdResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaMenu"}},"required":["ret"]},"GetApiV1MenustreeResp":{"title":"GetApiV1MenustreeResp","type":"object","properties":{"ret":{"type":"object"}},"required":["ret"]},"GetApiV1PubCurrentMenutreeResp":{"title":"GetApiV1PubCurrentMenutreeResp","type":"object","properties":{"ret":{"type":"object"}},"required":["ret"]},"GetApiV1PubCurrentUserResp":{"title":"GetApiV1PubCurrentUserResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaUserLoginInfo"}},"required":["ret"]},"GetApiV1PubLoginCaptchaResp":{"title":"GetApiV1PubLoginCaptchaResp","type":"object","properties":{"ret":{"type":"string"}},"required":["ret"]},"GetApiV1PubLoginCaptchaidResp":{"title":"GetApiV1PubLoginCaptchaidResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaLoginCaptcha"}},"required":["ret"]},"GetApiV1RolesResp":{"title":"GetApiV1RolesResp","type":"object","properties":{"ret":{"type":"object"}},"required":["ret"]},"GetApiV1Roles_IdResp":{"title":"GetApiV1Roles_IdResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaRole"}},"required":["ret"]},"GetApiV1RolesselectResp":{"title":"GetApiV1RolesselectResp","type":"object","properties":{"ret":{"type":"object"}},"required":["ret"]},"GetApiV1UsersResp":{"title":"GetApiV1UsersResp","type":"object","properties":{"ret":{"type":"object"}},"required":["ret"]},"GetApiV1Users_IdResp":{"title":"GetApiV1Users_IdResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaUser"}},"required":["ret"]},"PutApiV1Menus_IdResp":{"title":"PutApiV1Menus_IdResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaStatusResult"}},"required":["ret"]},"PutApiV1PubCurrentPasswordResp":{"title":"PutApiV1PubCurrentPasswordResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaStatusResult"}},"required":["ret"]},"PutApiV1Roles_IdResp":{"title":"PutApiV1Roles_IdResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaRole"}},"required":["ret"]},"PutApiV1Users_IdResp":{"title":"PutApiV1Users_IdResp","type":"object","properties":{"ret":{"$ref":"#/components/schemas/SchemaUser"}},"required":["ret"]},"SchemaErrorItem":{"title":"SchemaErrorItem","type":"object","properties":{"Code":{"type":"integer","format":"int32"},"Message":{"type":"string"}}},"SchemaErrorResult":{"title":"SchemaErrorResult","type":"object","properties":{"Error":{"$ref":"#/components/schemas/SchemaErrorItem"}}},"SchemaIDResult":{"title":"SchemaIDResult","type":"object","properties":{"Id":{"type":"string"}},"description":"\n\n\n"},"SchemaListResult":{"title":"SchemaListResult","type":"object","properties":{"List":{"type":"object"},"Pagination":{"$ref":"#/components/schemas/SchemaPaginationResult"}},"required":["List"]},"SchemaLoginCaptcha":{"title":"SchemaLoginCaptcha","type":"object","properties":{"CaptchaId":{"type":"string","description":"验证码ID"}},"description":"\n"},"SchemaLoginParam":{"title":"SchemaLoginParam","type":"object","properties":{"CaptchaCode":{"type":"string","description":"验证码\nrequired"},"CaptchaId":{"type":"string","description":"验证码ID\nrequired"},"Password":{"type":"string","description":"密码(md5加密)\nrequired"},"UserName":{"type":"string","description":"用户名\nrequired"}},"description":"\n","required":["CaptchaCode","CaptchaId","Password","UserName"]},"SchemaLoginTokenInfo":{"title":"SchemaLoginTokenInfo","type":"object","properties":{"AccessToken":{"type":"string","description":"访问令牌"},"ExpiresAt":{"type":"integer","format":"int32","description":"过期时间戳"},"TokenType":{"type":"string","description":"令牌类型"}},"description":"\n\n"},"SchemaMenu":{"title":"SchemaMenu","type":"object","properties":{"Actions":{"$ref":"#/components/schemas/SchemaMenuActions"},"CreatedAt":{"type":"string","description":"创建时间"},"Creator":{"type":"integer","format":"int32","description":"创建者"},"Icon":{"type":"string","description":"菜单图标"},"Id":{"type":"string","description":"唯一标识"},"IsShow":{"type":"integer","format":"int32","description":"是否显示(1:显示 2:隐藏)\nrequired"},"Memo":{"type":"string","description":"备注"},"Name":{"type":"string","description":"菜单名称\nrequired"},"ParentId":{"type":"string","description":"父级ID"},"ParentPath":{"type":"string","description":"父级路径"},"Router":{"type":"string","description":"访问路由"},"Sequence":{"type":"integer","format":"int32","description":"排序值"},"Status":{"type":"integer","format":"int32","description":"状态(1:启用 2:禁用)\nrequired"},"UpdatedAt":{"type":"string","description":"更新时间"}},"description":"\n\n\n","required":["IsShow","Name","Status"]},"SchemaMenuAction":{"title":"SchemaMenuAction","type":"object","properties":{"Code":{"type":"string","description":"动作编号\nrequired"},"Id":{"type":"string","description":"唯一标识"},"MenuId":{"type":"string","description":"菜单ID\nrequired"},"Name":{"type":"string","description":"动作名称\nrequired"},"Resources":{"$ref":"#/components/schemas/SchemaMenuActionResources"}},"required":["Code","MenuId","Name"]},"SchemaMenuActionResource":{"title":"SchemaMenuActionResource","type":"object","properties":{"ActionId":{"type":"string","description":"菜单动作ID"},"Id":{"type":"string","description":"唯一标识"},"Method":{"type":"string","description":"资源请求方式(支持正则)\nrequired"},"Path":{"type":"string","description":"资源请求路径（支持/:id匹配）\nrequired"}},"required":["Method","Path"]},"SchemaMenuActionResources":{"title":"SchemaMenuActionResources","type":"object"},"SchemaMenuActions":{"title":"SchemaMenuActions","type":"object"},"SchemaMenuTree":{"title":"SchemaMenuTree","type":"object","properties":{"Actions":{"$ref":"#/components/schemas/SchemaMenuActions"},"Children":{"$ref":"#/components/schemas/SchemaMenuTrees"},"Icon":{"type":"string","description":"菜单图标"},"Id":{"type":"string","description":"唯一标识"},"IsShow":{"type":"integer","format":"int32","description":"是否显示(1:显示 2:隐藏)"},"Name":{"type":"string","description":"菜单名称"},"ParentId":{"type":"string","description":"父级ID"},"ParentPath":{"type":"string","description":"父级路径"},"Router":{"type":"string","description":"访问路由"},"Sequence":{"type":"integer","format":"int32","description":"排序值"},"Status":{"type":"integer","format":"int32","description":"状态(1:启用 2:禁用)"}}},"SchemaMenuTrees":{"title":"SchemaMenuTrees","type":"object"},"SchemaPaginationResult":{"title":"SchemaPaginationResult","type":"object","properties":{"Current":{"type":"integer","format":"int32"},"PageSize":{"type":"integer","format":"int32"},"Total":{"type":"integer","format":"int32"}}},"SchemaRole":{"title":"SchemaRole","type":"object","properties":{"CreatedAt":{"type":"string","description":"创建时间"},"Creator":{"type":"integer","format":"int32","description":"创建者"},"Id":{"type":"string","description":"唯一标识"},"Memo":{"type":"string","description":"备注"},"Name":{"type":"string","description":"角色名称\nrequired"},"RoleMenus":{"$ref":"#/components/schemas/SchemaRoleMenus"},"Sequence":{"type":"integer","format":"int32","description":"排序值"},"Status":{"type":"integer","format":"int32","description":"状态(1:启用 2:禁用)\nrequired"},"UpdatedAt":{"type":"string","description":"更新时间"}},"description":"\n\n\n\n","required":["Name","RoleMenus","Status"]},"SchemaRoleMenu":{"title":"SchemaRoleMenu","type":"object","properties":{"ActionId":{"type":"string","description":"动作ID\nrequired"},"Id":{"type":"string","description":"唯一标识"},"MenuId":{"type":"string","description":"菜单ID\nrequired"},"RoleId":{"type":"string","description":"角色ID\nrequired"}},"required":["ActionId","MenuId","RoleId"]},"SchemaRoleMenus":{"title":"SchemaRoleMenus","type":"object"},"SchemaRoles":{"title":"SchemaRoles","type":"object"},"SchemaStatusResult":{"title":"SchemaStatusResult","type":"object","properties":{"Status":{"type":"string"}},"description":"\n\n\n\n\n\n"},"SchemaUpdatePasswordParam":{"title":"SchemaUpdatePasswordParam","type":"object","properties":{"NewPassword":{"type":"string","description":"新密码(md5加密)\nrequired"},"OldPassword":{"type":"string","description":"旧密码(md5加密)\nrequired"}},"description":"\n","required":["NewPassword","OldPassword"]},"SchemaUser":{"title":"SchemaUser","type":"object","properties":{"CreatedAt":{"type":"string","description":"创建时间"},"Creator":{"type":"integer","format":"int32","description":"创建者"},"Email":{"type":"string","description":"邮箱"},"Id":{"type":"string","description":"唯一标识"},"Password":{"type":"string","description":"密码"},"Phone":{"type":"string","description":"手机号"},"RealName":{"type":"string","description":"真实姓名\nrequired"},"Status":{"type":"integer","format":"int32","description":"用户状态(1:启用 2:停用)\nrequired"},"UserName":{"type":"string","description":"用户名\nrequired"},"UserRoles":{"$ref":"#/components/schemas/SchemaUserRoles"}},"description":"\n\n\n\n","required":["RealName","Status","UserName","UserRoles"]},"SchemaUserLoginInfo":{"title":"SchemaUserLoginInfo","type":"object","properties":{"RealName":{"type":"string","description":"真实姓名"},"Roles":{"$ref":"#/components/schemas/SchemaRoles"},"UserId":{"type":"string","description":"用户ID"},"UserName":{"type":"string","description":"用户名"}},"description":"\n"},"SchemaUserRole":{"title":"SchemaUserRole","type":"object","properties":{"Id":{"type":"string","description":"唯一标识"},"RoleId":{"type":"string","description":"角色ID"},"UserId":{"type":"string","description":"用户ID"}}},"SchemaUserRoles":{"title":"SchemaUserRoles","type":"object"},"SchemaUserShow":{"title":"SchemaUserShow","type":"object","properties":{"CreatedAt":{"type":"string","description":"创建时间"},"Email":{"type":"string","description":"邮箱"},"Id":{"type":"string","description":"唯一标识"},"Phone":{"type":"string","description":"手机号"},"RealName":{"type":"string","description":"真实姓名"},"Roles":{"type":"array","items":{"$ref":"#/components/schemas/SchemaRole"},"description":"授权角色列表"},"Status":{"type":"integer","format":"int32","description":"用户状态(1:启用 2:停用)"},"UserName":{"type":"string","description":"用户名"}},"required":["Roles"]}}}}`
}
