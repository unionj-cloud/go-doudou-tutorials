// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameQrtzBlobTrigger = "qrtz_blob_triggers"

// QrtzBlobTrigger mapped from table <qrtz_blob_triggers>
type QrtzBlobTrigger struct {
	SCHEDNAME    *string `gorm:"column:SCHED_NAME;type:varchar(120)" json:"SCHED_NAME,omitempty"`
	TRIGGERNAME  *string `gorm:"column:TRIGGER_NAME;type:varchar(200)" json:"TRIGGER_NAME,omitempty"`
	TRIGGERGROUP *string `gorm:"column:TRIGGER_GROUP;type:varchar(200)" json:"TRIGGER_GROUP,omitempty"`
	BLOBDATA     *[]byte `gorm:"column:BLOB_DATA;type:blob" json:"BLOB_DATA,omitempty"`
}

// TableName QrtzBlobTrigger's table name
func (*QrtzBlobTrigger) TableName() string {
	return TableNameQrtzBlobTrigger
}
