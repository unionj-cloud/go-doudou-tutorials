/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package dto

import (
	"time"
)

//go:generate go-doudou name --file $GOFILE

type Report struct {
	Id         string     `json:"id"`
	Code       *string    `json:"code"`
	Name       *string    `json:"name"`
	Note       *string    `json:"note"`
	Status     *string    `json:"status"`
	Type       *string    `json:"type"`
	JsonStr    *string    `json:"jsonStr"`
	ApiUrl     *string    `json:"apiUrl"`
	Thumb      *string    `json:"thumb"`
	CreateBy   *string    `json:"createBy"`
	CreateTime *time.Time `json:"createTime"`
	UpdateBy   *string    `json:"updateBy"`
	UpdateTime *time.Time `json:"updateTime"`
	DelFlag    *int8      `json:"delFlag"`
	ApiMethod  *string    `json:"apiMethod"`
	ApiCode    *string    `json:"apiCode"`
	Template   *int8      `json:"template"`
	ViewCount  *int64     `json:"viewCount"`
}

type ReportDataSource struct {
	Id           string     `json:"id"`
	Name         *string    `json:"name"`
	ReportId     *string    `json:"reportId"`
	Code         *string    `json:"code"`
	Remark       *string    `json:"remark"`
	DbType       *string    `json:"dbType"`
	DbDriver     *string    `json:"dbDriver"`
	DbUrl        *string    `json:"dbUrl"`
	DbUsername   *string    `json:"dbUsername"`
	DbPassword   *string    `json:"dbPassword"`
	CreateBy     *string    `json:"createBy"`
	CreateTime   *time.Time `json:"createTime"`
	UpdateBy     *string    `json:"updateBy"`
	UpdateTime   *time.Time `json:"updateTime"`
	ConnectTimes *int       `json:"connectTimes"`
}

type ReportDb struct {
	Id            string     `json:"id"`
	JimuReportId  *string    `json:"jimuReportId"`
	CreateBy      *string    `json:"createBy"`
	UpdateBy      *string    `json:"updateBy"`
	CreateTime    *time.Time `json:"createTime"`
	UpdateTime    *time.Time `json:"updateTime"`
	DbCode        *string    `json:"dbCode"`
	DbChName      *string    `json:"dbChName"`
	DbType        *string    `json:"dbType"`
	DbTableName   *string    `json:"dbTableName"`
	DbDynSql      *string    `json:"dbDynSql"`
	DbKey         *string    `json:"dbKey"`
	TbDbKey       *string    `json:"tbDbKey"`
	TbDbTableName *string    `json:"tbDbTableName"`
	JavaType      *string    `json:"javaType"`
	JavaValue     *string    `json:"javaValue"`
	ApiUrl        *string    `json:"apiUrl"`
	ApiMethod     *string    `json:"apiMethod"`
	IsList        *int       `json:"isList"`
	IsPage        *string    `json:"isPage"`
	DbSource      *string    `json:"dbSource"`
	DbSourceType  *string    `json:"dbSourceType"`
	JsonData      *string    `json:"jsonData"`
	ApiConvert    *string    `json:"apiConvert"`
}

type ReportDbField struct {
	Id             string     `json:"id"`
	CreateBy       *string    `json:"createBy"`
	CreateTime     *time.Time `json:"createTime"`
	UpdateBy       *string    `json:"updateBy"`
	UpdateTime     *time.Time `json:"updateTime"`
	JimuReportDbId *string    `json:"jimuReportDbId"`
	FieldName      *string    `json:"fieldName"`
	FieldText      *string    `json:"fieldText"`
	WidgetType     *string    `json:"widgetType"`
	WidgetWidth    *int       `json:"widgetWidth"`
	OrderNum       *int       `json:"orderNum"`
	SearchFlag     *int       `json:"searchFlag"`
	SearchMode     *int       `json:"searchMode"`
	DictCode       *string    `json:"dictCode"`
	SearchValue    *string    `json:"searchValue"`
}

type ReportDbParam struct {
	Id               string     `json:"id"`
	JimuReportHeadId string     `json:"jimuReportHeadId"`
	ParamName        string     `json:"paramName"`
	ParamTxt         *string    `json:"paramTxt"`
	ParamValue       *string    `json:"paramValue"`
	OrderNum         *int       `json:"orderNum"`
	CreateBy         *string    `json:"createBy"`
	CreateTime       *time.Time `json:"createTime"`
	UpdateBy         *string    `json:"updateBy"`
	UpdateTime       *time.Time `json:"updateTime"`
}

type ReportLink struct {
	Id          string  `json:"id"`
	ReportId    *string `json:"reportId"`
	Parameter   *string `json:"parameter"`
	EjectType   *string `json:"ejectType"`
	LinkName    *string `json:"linkName"`
	ApiMethod   *string `json:"apiMethod"`
	LinkType    *string `json:"linkType"`
	ApiUrl      *string `json:"apiUrl"`
	LinkChartId *string `json:"linkChartId"`
}

type ReportMap struct {
	Id         string     `json:"id"`
	Label      *string    `json:"label"`
	Name       *string    `json:"name"`
	Data       *string    `json:"data"`
	CreateBy   *string    `json:"createBy"`
	CreateTime *time.Time `json:"createTime"`
	UpdateBy   *string    `json:"updateBy"`
	UpdateTime *time.Time `json:"updateTime"`
	DelFlag    *string    `json:"delFlag"`
	SysOrgCode *string    `json:"sysOrgCode"`
}

type ReportShare struct {
	Id             string     `json:"id"`
	ReportId       *string    `json:"reportId"`
	PreviewUrl     *string    `json:"previewUrl"`
	PreviewLock    *string    `json:"previewLock"`
	LastUpdateTime *time.Time `json:"lastUpdateTime"`
	TermOfValidity *string    `json:"termOfValidity"`
	Status         *string    `json:"status"`
}

type CommonComment struct {
	Id                  string     `json:"id"`
	DelFlag             *int8      `json:"delFlag"`
	RelateBusinessTable *string    `json:"relateBusinessTable"`
	RelateBusinessId    *string    `json:"relateBusinessId"`
	Content             *string    `json:"content"`
	ReplyTo             *string    `json:"replyTo"`
	CreateTime          *time.Time `json:"createTime"`
	CreateBy            *string    `json:"createBy"`
	SysOrgCode          *string    `json:"sysOrgCode"`
}

type CommonConfig struct {
	Id           string     `json:"id"`
	DelFlag      *int8      `json:"delFlag"`
	Name         *string    `json:"name"`
	Code         *string    `json:"code"`
	Value        *string    `json:"value"`
	DefaultValue *string    `json:"defaultValue"`
	ConfigDesc   *string    `json:"configDesc"`
	CreateTime   *time.Time `json:"createTime"`
	CreateBy     *string    `json:"createBy"`
	UpdateTime   *time.Time `json:"updateTime"`
	UpdateBy     *string    `json:"updateBy"`
	SysOrgCode   *string    `json:"sysOrgCode"`
}

type CommonFile struct {
	Id                  string     `json:"id"`
	DelFlag             *int8      `json:"delFlag"`
	RelateBusinessTable *string    `json:"relateBusinessTable"`
	RelateBusinessId    *string    `json:"relateBusinessId"`
	FileName            *string    `json:"fileName"`
	FileType            *string    `json:"fileType"`
	FileSize            *int64     `json:"fileSize"`
	FilePath            *string    `json:"filePath"`
	CreateTime          *time.Time `json:"createTime"`
	CreateBy            *string    `json:"createBy"`
	SysOrgCode          *string    `json:"sysOrgCode"`
}

type CommonFollow struct {
	Id                  string     `json:"id"`
	RelateBusinessTable *string    `json:"relateBusinessTable"`
	RelateBusinessId    *string    `json:"relateBusinessId"`
	CreateTime          *time.Time `json:"createTime"`
	CreateBy            *string    `json:"createBy"`
	SysOrgCode          *string    `json:"sysOrgCode"`
}

type CommonLike struct {
	Id                  string     `json:"id"`
	RelateBusinessTable *string    `json:"relateBusinessTable"`
	RelateBusinessId    *string    `json:"relateBusinessId"`
	CreateTime          *time.Time `json:"createTime"`
	CreateBy            *string    `json:"createBy"`
	SysOrgCode          *string    `json:"sysOrgCode"`
}

type CommonLog struct {
	Id                  string     `json:"id"`
	RelateBusinessTable *string    `json:"relateBusinessTable"`
	RelateBusinessId    *string    `json:"relateBusinessId"`
	Type                *string    `json:"type"`
	SubType             *string    `json:"subType"`
	Desc                *string    `json:"desc"`
	CreateTime          *time.Time `json:"createTime"`
	CreateBy            *string    `json:"createBy"`
	SysOrgCode          *string    `json:"sysOrgCode"`
}

type CommonUser struct {
	Id               string  `json:"id"`
	Nickname         *string `json:"nickname"`
	RealName         *string `json:"realName"`
	Organization     *string `json:"organization"`
	OrganizationCode *string `json:"organizationCode"`
	Company          *string `json:"company"`
	CompanyCode      *string `json:"companyCode"`
	Identity         *string `json:"identity"`
	IdentityCode     *string `json:"identityCode"`
	Industry         *string `json:"industry"`
	IndustryCode     *string `json:"industryCode"`
	Profession       *string `json:"profession"`
	ProfessionCode   *string `json:"professionCode"`
	CadPurpose       *string `json:"cadPurpose"`
	Score            *int    `json:"score"`
	WxAccount        *string `json:"wxAccount"`
	ContactInfo      *string `json:"contactInfo"`
	OrganizationName *string `json:"organizationName"`
}

type CommonUserFile struct {
	Id           int        `json:"id"`
	UserId       *string    `json:"userId"`
	FileCategory *int       `json:"fileCategory"`
	FilePath     *string    `json:"filePath"`
	CreateTime   *time.Time `json:"createTime"`
}

type CommonUserOrgDict struct {
	Id         string     `json:"id"`
	Pid        *string    `json:"pid"`
	Name       *string    `json:"name"`
	Type       *string    `json:"type"`
	DelFlag    *int8      `json:"delFlag"`
	CreateBy   *string    `json:"createBy"`
	CreateTime *time.Time `json:"createTime"`
	UpdateBy   *string    `json:"updateBy"`
	UpdateTime *time.Time `json:"updateTime"`
}

type Knowledge struct {
	Id                   string     `json:"id"`
	DelFlag              *int8      `json:"delFlag"`
	Title                *string    `json:"title"`
	Content              *string    `json:"content"`
	ViewCount            *int       `json:"viewCount"`
	Catalog              *string    `json:"catalog"`
	Tag                  *string    `json:"tag"`
	Status               *string    `json:"status"`
	LastReleaseVersionId *string    `json:"lastReleaseVersionId"`
	LastReleaseTime      *time.Time `json:"lastReleaseTime"`
	LastReleaseBy        *string    `json:"lastReleaseBy"`
	CreateTime           *time.Time `json:"createTime"`
	CreateBy             *string    `json:"createBy"`
	UpdateTime           *time.Time `json:"updateTime"`
	UpdateBy             *string    `json:"updateBy"`
	SysOrgCode           *string    `json:"sysOrgCode"`
	ReferenceDocuments   *string    `json:"referenceDocuments"`
	SortOrder            *int       `json:"sortOrder"`
	ActualTitles         *string    `json:"actualTitles"`
}

type KnowledgeCatalog struct {
	Id         string     `json:"id"`
	Pid        *string    `json:"pid"`
	Name       *string    `json:"name"`
	Code       *string    `json:"code"`
	SortOrder  *int       `json:"sortOrder"`
	DelFlag    *int8      `json:"delFlag"`
	CreateBy   *string    `json:"createBy"`
	CreateTime *time.Time `json:"createTime"`
	UpdateBy   *string    `json:"updateBy"`
	UpdateTime *time.Time `json:"updateTime"`
	HasChild   *string    `json:"hasChild"`
}

type KnowledgeVersion struct {
	VersionId          string     `json:"versionId"`
	DelFlag            *int8      `json:"delFlag"`
	Title              *string    `json:"title"`
	Content            *string    `json:"content"`
	ViewCount          *int       `json:"viewCount"`
	Catalog            *string    `json:"catalog"`
	Tag                *string    `json:"tag"`
	VersionStatus      *string    `json:"versionStatus"`
	EditReason         *string    `json:"editReason"`
	EditRemark         *string    `json:"editRemark"`
	CensorBy           *string    `json:"censorBy"`
	CensorTime         *time.Time `json:"censorTime"`
	RejectReason       *string    `json:"rejectReason"`
	CreateTime         *time.Time `json:"createTime"`
	CreateBy           *string    `json:"createBy"`
	UpdateTime         *time.Time `json:"updateTime"`
	UpdateBy           *string    `json:"updateBy"`
	SysOrgCode         *string    `json:"sysOrgCode"`
	ReferenceDocuments *string    `json:"referenceDocuments"`
	CadKnowledgeId     *string    `json:"cadKnowledgeId"`
}

type SquareCommit struct {
	Id          string     `json:"id"`
	DelFlag     *int8      `json:"delFlag"`
	TaskId      *string    `json:"taskId"`
	Version     *string    `json:"version"`
	Input       *string    `json:"input"`
	Output      *string    `json:"output"`
	IsBest      *int8      `json:"isBest"`
	IsPublished *int8      `json:"isPublished"`
	Remark      *string    `json:"remark"`
	CreateTime  *time.Time `json:"createTime"`
	CreateBy    *string    `json:"createBy"`
	UpdateTime  *time.Time `json:"updateTime"`
	UpdateBy    *string    `json:"updateBy"`
	SysOrgCode  *string    `json:"sysOrgCode"`
	AuditStatus *int8      `json:"auditStatus"`
	IsRead      *int8      `json:"isRead"`
	FilePath    *string    `json:"filePath"`
}

type SquareTask struct {
	Id                 string     `json:"id"`
	DelFlag            *int8      `json:"delFlag"`
	Status             *string    `json:"status"`
	Name               *string    `json:"name"`
	Difficulty         *string    `json:"difficulty"`
	EffectiveTimeStart *time.Time `json:"effectiveTimeStart"`
	EffectiveTimeEnd   *time.Time `json:"effectiveTimeEnd"`
	Cover              *string    `json:"cover"`
	InputExample       *string    `json:"inputExample"`
	OutputExample      *string    `json:"outputExample"`
	Tag                *string    `json:"tag"`
	ViewCount          *int       `json:"viewCount"`
	PublishTime        *time.Time `json:"publishTime"`
	CreateTime         *time.Time `json:"createTime"`
	CreateBy           *string    `json:"createBy"`
	UpdateTime         *time.Time `json:"updateTime"`
	UpdateBy           *string    `json:"updateBy"`
	SysOrgCode         *string    `json:"sysOrgCode"`
	Priority           *string    `json:"priority"`
	EditSnapshot       *string    `json:"editSnapshot"`
}

type SquareTaskAudit struct {
	Id          string     `json:"id"`
	TaskId      *string    `json:"taskId"`
	AuditStatus *string    `json:"auditStatus"`
	Reason      *string    `json:"reason"`
	CreateTime  *time.Time `json:"createTime"`
	CreateBy    *string    `json:"createBy"`
	CommitTime  *time.Time `json:"commitTime"`
	UpdateTime  *time.Time `json:"updateTime"`
	UpdateBy    *string    `json:"updateBy"`
}

type SquareTaskKnowledge struct {
	Id          string     `json:"id"`
	TaskId      *string    `json:"taskId"`
	KnowledgeId *string    `json:"knowledgeId"`
	CreateTime  *time.Time `json:"createTime"`
}

type SquareTaskReject struct {
	Id         string     `json:"id"`
	TaskId     *string    `json:"taskId"`
	Reason     *string    `json:"reason"`
	CreateTime *time.Time `json:"createTime"`
	CreateBy   *string    `json:"createBy"`
}

type SquareTaskUser struct {
	Id         string     `json:"id"`
	IsAssigned *int8      `json:"isAssigned"`
	TaskId     *string    `json:"taskId"`
	UserId     *string    `json:"userId"`
	CreateTime *time.Time `json:"createTime"`
}

type FieldKongj struct {
	Id         string
	CreateBy   *string
	CreateTime *time.Time
	UpdateBy   *string
	UpdateTime *time.Time
	SysOrgCode *string
	Name       *string
	Sex        *string
	Radio      *string
	Checkbox   *string
	SelMut     *string
	SelSearch  *string
	Birthday   *time.Time
	Pic        *string
	Files      *string
	Remakr     *string
	Fuwenb     *string
	UserSel    *string
	DepSel     *string
	Ddd        *float64
}

type OrderCustomer struct {
	Id          string
	CreateBy    *string
	CreateTime  *time.Time
	UpdateBy    *string
	UpdateTime  *time.Time
	SysOrgCode  *string
	Name        *string
	Sex         *string
	Birthday    *time.Time
	Age         *int
	Address     *string
	OrderMainId *string
}

type OrderGoods struct {
	Id          string
	CreateBy    *string
	CreateTime  *time.Time
	UpdateBy    *string
	UpdateTime  *time.Time
	SysOrgCode  *string
	GoodName    *string
	Price       *float64
	Num         *int
	ZongPrice   *float64
	OrderMainId *string
}

type OrderMain struct {
	Id         string
	CreateBy   *string
	CreateTime *time.Time
	UpdateBy   *string
	UpdateTime *time.Time
	SysOrgCode *string
	OrderCode  *string
	XdDate     *time.Time
	Money      *float64
	Remark     *string
}

type ShopType struct {
	Id         string
	CreateBy   *string
	CreateTime *time.Time
	UpdateBy   *string
	UpdateTime *time.Time
	SysOrgCode *string
	Name       *string
	Content    *string
	Pics       *string
	Pid        *string
	HasChild   *string
}
