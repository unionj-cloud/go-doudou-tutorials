// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameQrtzBlobTrigger = "qrtz_blob_triggers"

// QrtzBlobTrigger mapped from table <qrtz_blob_triggers>
type QrtzBlobTrigger struct {
	SCHEDNAME    string `gorm:"column:SCHED_NAME" json:"SCHED_NAME"`
	TRIGGERNAME  string `gorm:"column:TRIGGER_NAME" json:"TRIGGER_NAME"`
	TRIGGERGROUP string `gorm:"column:TRIGGER_GROUP" json:"TRIGGER_GROUP"`
	BLOBDATA     []byte `gorm:"column:BLOB_DATA" json:"BLOB_DATA"`
}

// TableName QrtzBlobTrigger's table name
func (*QrtzBlobTrigger) TableName() string {
	return TableNameQrtzBlobTrigger
}
