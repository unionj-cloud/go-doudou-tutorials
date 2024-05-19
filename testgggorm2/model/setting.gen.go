// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"fmt"
	"testgggorm2/config"

	"github.com/unionj-cloud/go-doudou/v2/toolkit/stringutils"
)

var TableNameSetting string

func init() {
	if stringutils.IsNotEmpty(config.G_Config.Db.Name) {
		TableNameSetting = fmt.Sprintf("%s.setting", config.G_Config.Db.Name)
	} else {
		TableNameSetting = "setting"
	}
}

// Setting mapped from table <setting>
type Setting struct {
	ID              int32   `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,omitempty"`
	SiteTitle       *string `gorm:"column:site_title;type:varchar(100);comment:商城名称" json:"siteTitle,omitempty"`             // 商城名称
	SiteLogo        *string `gorm:"column:site_logo;type:varchar(250);comment:商城图标" json:"siteLogo,omitempty"`               // 商城图标
	SiteKeywords    *string `gorm:"column:site_keywords;type:varchar(100);comment:商城关键字" json:"siteKeywords,omitempty"`      // 商城关键字
	SiteDescription *string `gorm:"column:site_description;type:varchar(500);comment:商城描述" json:"siteDescription,omitempty"` // 商城描述
	NoPicture       *string `gorm:"column:no_picture;type:varchar(100);comment:没有图片显示" json:"noPicture,omitempty"`           // 没有图片显示
	SiteIcp         *string `gorm:"column:site_icp;type:varchar(50);comment:商城ICP" json:"siteIcp,omitempty"`                 // 商城ICP
	SiteTel         *string `gorm:"column:site_tel;type:varchar(50);comment:商城手机号" json:"siteTel,omitempty"`                 // 商城手机号
	SearchKeywords  *string `gorm:"column:search_keywords;type:varchar(250);comment:搜索关键字" json:"searchKeywords,omitempty"`  // 搜索关键字
	TongjiCode      *string `gorm:"column:tongji_code;type:varchar(500);comment:统计编码" json:"tongjiCode,omitempty"`           // 统计编码
	Appid           *string `gorm:"column:appid;type:varchar(50);comment:oss appid" json:"appid,omitempty"`                  // oss appid
	AppSecret       *string `gorm:"column:app_secret;type:varchar(80);comment:oss app_secret" json:"appSecret,omitempty"`    // oss app_secret
	EndPoint        *string `gorm:"column:end_point;type:varchar(200);comment:oss 终端点" json:"endPoint,omitempty"`            // oss 终端点
	BucketName      *string `gorm:"column:bucket_name;type:varchar(200);comment:oss 桶名称" json:"bucketName,omitempty"`        // oss 桶名称
	OssStatus       *int32  `gorm:"column:oss_status;type:tinyint(4);comment:oss 状态" json:"ossStatus,omitempty"`             // oss 状态
}

// TableName Setting's table name
func (*Setting) TableName() string {
	return TableNameSetting
}
