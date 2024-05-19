// Code generated by gorm.io/gen for go-doudou. YOU CAN EDIT.
// Code generated by gorm.io/gen for go-doudou. YOU CAN EDIT.
// Code generated by gorm.io/gen for go-doudou. YOU CAN EDIT.

package service

import (
	"context"
	"testgggorm2/dto"
)

//go:generate go-doudou svc http -c
//go:generate go-doudou svc grpc

type Testgggorm2 interface {

	// PostGenMenu mapped from table <menu>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenMenu(ctx context.Context, body dto.Menu) (data int32, err error)

	// PostGenMenus mapped from table <menu>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenMenus(ctx context.Context, body []dto.Menu) (data []int32, err error)

	// GetGenMenu_Id mapped from table <menu>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenMenu_Id(ctx context.Context, id int32) (data dto.Menu, err error)

	// PutGenMenu mapped from table <menu>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenMenu(ctx context.Context, body dto.Menu) error

	// DeleteGenMenu_Id mapped from table <menu>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenMenu_Id(ctx context.Context, id int32) error

	// GetGenMenus mapped from table <menu>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenMenus(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenProduct mapped from table <product>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenProduct(ctx context.Context, body dto.Product) (data int32, err error)

	// PostGenProducts mapped from table <product>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenProducts(ctx context.Context, body []dto.Product) (data []int32, err error)

	// GetGenProduct_Id mapped from table <product>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenProduct_Id(ctx context.Context, id int32) (data dto.Product, err error)

	// PutGenProduct mapped from table <product>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenProduct(ctx context.Context, body dto.Product) error

	// DeleteGenProduct_Id mapped from table <product>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenProduct_Id(ctx context.Context, id int32) error

	// GetGenProducts mapped from table <product>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenProducts(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenProductAttr mapped from table <product_attr>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenProductAttr(ctx context.Context, body dto.ProductAttr) (data int32, err error)

	// PostGenProductAttrs mapped from table <product_attr>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenProductAttrs(ctx context.Context, body []dto.ProductAttr) (data []int32, err error)

	// GetGenProductAttr_Id mapped from table <product_attr>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenProductAttr_Id(ctx context.Context, id int32) (data dto.ProductAttr, err error)

	// PutGenProductAttr mapped from table <product_attr>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenProductAttr(ctx context.Context, body dto.ProductAttr) error

	// DeleteGenProductAttr_Id mapped from table <product_attr>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenProductAttr_Id(ctx context.Context, id int32) error

	// GetGenProductAttrs mapped from table <product_attr>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenProductAttrs(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenProductImage mapped from table <product_image>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenProductImage(ctx context.Context, body dto.ProductImage) (data int32, err error)

	// PostGenProductImages mapped from table <product_image>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenProductImages(ctx context.Context, body []dto.ProductImage) (data []int32, err error)

	// GetGenProductImage_Id mapped from table <product_image>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenProductImage_Id(ctx context.Context, id int32) (data dto.ProductImage, err error)

	// PutGenProductImage mapped from table <product_image>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenProductImage(ctx context.Context, body dto.ProductImage) error

	// DeleteGenProductImage_Id mapped from table <product_image>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenProductImage_Id(ctx context.Context, id int32) error

	// GetGenProductImages mapped from table <product_image>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenProductImages(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenRole mapped from table <role>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenRole(ctx context.Context, body dto.Role) (data int32, err error)

	// PostGenRoles mapped from table <role>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenRoles(ctx context.Context, body []dto.Role) (data []int32, err error)

	// GetGenRole_Id mapped from table <role>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenRole_Id(ctx context.Context, id int32) (data dto.Role, err error)

	// PutGenRole mapped from table <role>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenRole(ctx context.Context, body dto.Role) error

	// DeleteGenRole_Id mapped from table <role>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenRole_Id(ctx context.Context, id int32) error

	// GetGenRoles mapped from table <role>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenRoles(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenRoleAuth mapped from table <role_auth>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenRoleAuth(ctx context.Context, body dto.RoleAuth) (data int32, err error)

	// PostGenRoleAuths mapped from table <role_auth>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenRoleAuths(ctx context.Context, body []dto.RoleAuth) (data []int32, err error)

	// GetGenRoleAuth_Id mapped from table <role_auth>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenRoleAuth_Id(ctx context.Context, id int32) (data dto.RoleAuth, err error)

	// PutGenRoleAuth mapped from table <role_auth>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenRoleAuth(ctx context.Context, body dto.RoleAuth) error

	// DeleteGenRoleAuth_Id mapped from table <role_auth>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenRoleAuth_Id(ctx context.Context, id int32) error

	// GetGenRoleAuths mapped from table <role_auth>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenRoleAuths(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenSetting mapped from table <setting>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenSetting(ctx context.Context, body dto.Setting) (data int32, err error)

	// PostGenSettings mapped from table <setting>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenSettings(ctx context.Context, body []dto.Setting) (data []int32, err error)

	// GetGenSetting_Id mapped from table <setting>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenSetting_Id(ctx context.Context, id int32) (data dto.Setting, err error)

	// PutGenSetting mapped from table <setting>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenSetting(ctx context.Context, body dto.Setting) error

	// DeleteGenSetting_Id mapped from table <setting>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenSetting_Id(ctx context.Context, id int32) error

	// GetGenSettings mapped from table <setting>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenSettings(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenAuth mapped from table <auth>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenAuth(ctx context.Context, body dto.Auth) (data int32, err error)

	// PostGenAuths mapped from table <auth>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenAuths(ctx context.Context, body []dto.Auth) (data []int32, err error)

	// GetGenAuth_Id mapped from table <auth>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenAuth_Id(ctx context.Context, id int32) (data dto.Auth, err error)

	// PutGenAuth mapped from table <auth>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenAuth(ctx context.Context, body dto.Auth) error

	// DeleteGenAuth_Id mapped from table <auth>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenAuth_Id(ctx context.Context, id int32) error

	// GetGenAuths mapped from table <auth>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenAuths(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenProductType mapped from table <product_type>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenProductType(ctx context.Context, body dto.ProductType) (data int32, err error)

	// PostGenProductTypes mapped from table <product_type>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenProductTypes(ctx context.Context, body []dto.ProductType) (data []int32, err error)

	// GetGenProductType_Id mapped from table <product_type>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenProductType_Id(ctx context.Context, id int32) (data dto.ProductType, err error)

	// PutGenProductType mapped from table <product_type>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenProductType(ctx context.Context, body dto.ProductType) error

	// DeleteGenProductType_Id mapped from table <product_type>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenProductType_Id(ctx context.Context, id int32) error

	// GetGenProductTypes mapped from table <product_type>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenProductTypes(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenProductTypeAttribute mapped from table <product_type_attribute>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenProductTypeAttribute(ctx context.Context, body dto.ProductTypeAttribute) (data int32, err error)

	// PostGenProductTypeAttributes mapped from table <product_type_attribute>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenProductTypeAttributes(ctx context.Context, body []dto.ProductTypeAttribute) (data []int32, err error)

	// GetGenProductTypeAttribute_Id mapped from table <product_type_attribute>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenProductTypeAttribute_Id(ctx context.Context, id int32) (data dto.ProductTypeAttribute, err error)

	// PutGenProductTypeAttribute mapped from table <product_type_attribute>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenProductTypeAttribute(ctx context.Context, body dto.ProductTypeAttribute) error

	// DeleteGenProductTypeAttribute_Id mapped from table <product_type_attribute>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenProductTypeAttribute_Id(ctx context.Context, id int32) error

	// GetGenProductTypeAttributes mapped from table <product_type_attribute>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenProductTypeAttributes(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenSchemaTest mapped from table <schema_test>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenSchemaTest(ctx context.Context, body dto.SchemaTest) (data int32, err error)

	// PostGenSchemaTests mapped from table <schema_test>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenSchemaTests(ctx context.Context, body []dto.SchemaTest) (data []int32, err error)

	// GetGenSchemaTest_Id mapped from table <schema_test>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenSchemaTest_Id(ctx context.Context, id int32) (data dto.SchemaTest, err error)

	// PutGenSchemaTest mapped from table <schema_test>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenSchemaTest(ctx context.Context, body dto.SchemaTest) error

	// DeleteGenSchemaTest_Id mapped from table <schema_test>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenSchemaTest_Id(ctx context.Context, id int32) error

	// GetGenSchemaTests mapped from table <schema_test>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenSchemaTests(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenUser mapped from table <user>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenUser(ctx context.Context, body dto.User) (data int32, err error)

	// PostGenUsers mapped from table <user>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenUsers(ctx context.Context, body []dto.User) (data []int32, err error)

	// GetGenUser_Id mapped from table <user>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenUser_Id(ctx context.Context, id int32) (data dto.User, err error)

	// PutGenUser mapped from table <user>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenUser(ctx context.Context, body dto.User) error

	// DeleteGenUser_Id mapped from table <user>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenUser_Id(ctx context.Context, id int32) error

	// GetGenUsers mapped from table <user>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenUsers(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenUserSm mapped from table <user_sms>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenUserSm(ctx context.Context, body dto.UserSm) (data int32, err error)

	// PostGenUserSms mapped from table <user_sms>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenUserSms(ctx context.Context, body []dto.UserSm) (data []int32, err error)

	// GetGenUserSm_Id mapped from table <user_sms>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenUserSm_Id(ctx context.Context, id int32) (data dto.UserSm, err error)

	// PutGenUserSm mapped from table <user_sms>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenUserSm(ctx context.Context, body dto.UserSm) error

	// DeleteGenUserSm_Id mapped from table <user_sms>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenUserSm_Id(ctx context.Context, id int32) error

	// GetGenUserSms mapped from table <user_sms>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenUserSms(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenBanner mapped from table <banner>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenBanner(ctx context.Context, body dto.Banner) (data int32, err error)

	// PostGenBanners mapped from table <banner>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenBanners(ctx context.Context, body []dto.Banner) (data []int32, err error)

	// GetGenBanner_Id mapped from table <banner>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenBanner_Id(ctx context.Context, id int32) (data dto.Banner, err error)

	// PutGenBanner mapped from table <banner>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenBanner(ctx context.Context, body dto.Banner) error

	// DeleteGenBanner_Id mapped from table <banner>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenBanner_Id(ctx context.Context, id int32) error

	// GetGenBanners mapped from table <banner>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenBanners(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenCart mapped from table <cart>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenCart(ctx context.Context, body dto.Cart) (data int32, err error)

	// PostGenCarts mapped from table <cart>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenCarts(ctx context.Context, body []dto.Cart) (data []int32, err error)

	// GetGenCart_Id mapped from table <cart>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenCart_Id(ctx context.Context, id int32) (data dto.Cart, err error)

	// PutGenCart mapped from table <cart>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenCart(ctx context.Context, body dto.Cart) error

	// DeleteGenCart_Id mapped from table <cart>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenCart_Id(ctx context.Context, id int32) error

	// GetGenCarts mapped from table <cart>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenCarts(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenOrderItem mapped from table <order_item>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenOrderItem(ctx context.Context, body dto.OrderItem) (data int32, err error)

	// PostGenOrderItems mapped from table <order_item>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenOrderItems(ctx context.Context, body []dto.OrderItem) (data []int32, err error)

	// GetGenOrderItem_Id mapped from table <order_item>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenOrderItem_Id(ctx context.Context, id int32) (data dto.OrderItem, err error)

	// PutGenOrderItem mapped from table <order_item>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenOrderItem(ctx context.Context, body dto.OrderItem) error

	// DeleteGenOrderItem_Id mapped from table <order_item>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenOrderItem_Id(ctx context.Context, id int32) error

	// GetGenOrderItems mapped from table <order_item>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenOrderItems(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenAddress mapped from table <address>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenAddress(ctx context.Context, body dto.Address) (data int32, err error)

	// PostGenAddresss mapped from table <address>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenAddresss(ctx context.Context, body []dto.Address) (data []int32, err error)

	// GetGenAddress_Id mapped from table <address>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenAddress_Id(ctx context.Context, id int32) (data dto.Address, err error)

	// PutGenAddress mapped from table <address>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenAddress(ctx context.Context, body dto.Address) error

	// DeleteGenAddress_Id mapped from table <address>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenAddress_Id(ctx context.Context, id int32) error

	// GetGenAddresss mapped from table <address>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenAddresss(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenAdministrator mapped from table <administrator>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenAdministrator(ctx context.Context, body dto.Administrator) (data int32, err error)

	// PostGenAdministrators mapped from table <administrator>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenAdministrators(ctx context.Context, body []dto.Administrator) (data []int32, err error)

	// GetGenAdministrator_Id mapped from table <administrator>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenAdministrator_Id(ctx context.Context, id int32) (data dto.Administrator, err error)

	// PutGenAdministrator mapped from table <administrator>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenAdministrator(ctx context.Context, body dto.Administrator) error

	// DeleteGenAdministrator_Id mapped from table <administrator>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenAdministrator_Id(ctx context.Context, id int32) error

	// GetGenAdministrators mapped from table <administrator>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenAdministrators(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenOrder mapped from table <order>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenOrder(ctx context.Context, body dto.Order) (data int32, err error)

	// PostGenOrders mapped from table <order>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenOrders(ctx context.Context, body []dto.Order) (data []int32, err error)

	// GetGenOrder_Id mapped from table <order>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenOrder_Id(ctx context.Context, id int32) (data dto.Order, err error)

	// PutGenOrder mapped from table <order>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenOrder(ctx context.Context, body dto.Order) error

	// DeleteGenOrder_Id mapped from table <order>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenOrder_Id(ctx context.Context, id int32) error

	// GetGenOrders mapped from table <order>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenOrders(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenProductCate mapped from table <product_cate>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenProductCate(ctx context.Context, body dto.ProductCate) (data int32, err error)

	// PostGenProductCates mapped from table <product_cate>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenProductCates(ctx context.Context, body []dto.ProductCate) (data []int32, err error)

	// GetGenProductCate_Id mapped from table <product_cate>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenProductCate_Id(ctx context.Context, id int32) (data dto.ProductCate, err error)

	// PutGenProductCate mapped from table <product_cate>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenProductCate(ctx context.Context, body dto.ProductCate) error

	// DeleteGenProductCate_Id mapped from table <product_cate>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenProductCate_Id(ctx context.Context, id int32) error

	// GetGenProductCates mapped from table <product_cate>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenProductCates(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)

	// PostGenProductColor mapped from table <product_color>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenProductColor(ctx context.Context, body dto.ProductColor) (data int32, err error)

	// PostGenProductColors mapped from table <product_color>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PostGenProductColors(ctx context.Context, body []dto.ProductColor) (data []int32, err error)

	// GetGenProductColor_Id mapped from table <product_color>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenProductColor_Id(ctx context.Context, id int32) (data dto.ProductColor, err error)

	// PutGenProductColor mapped from table <product_color>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	PutGenProductColor(ctx context.Context, body dto.ProductColor) error

	// DeleteGenProductColor_Id mapped from table <product_color>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	DeleteGenProductColor_Id(ctx context.Context, id int32) error

	// GetGenProductColors mapped from table <product_color>
	// Code generated by gorm.io/gen for go-doudou. DO NOT EDIT.
	GetGenProductColors(ctx context.Context, parameter dto.Parameter) (data dto.Page, err error)
}