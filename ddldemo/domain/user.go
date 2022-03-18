package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

//dd:table
type User struct {
	Id        int             `dd:"pk;auto;type:int(10) unsigned"`
	Name      string          `dd:"type:varchar(45);index:name_phone_idx,1,asc"`
	Phone     string          `dd:"type:varchar(45);index:name_phone_idx,2,asc"`
	Age       int             `dd:"type:int(11)"`
	No        int             `dd:"type:int(11)"`
	School    *string         `dd:"type:varchar(255);default:'harvard';extra:comment 'school';index:school_idx,1,asc"`
	IsStudent int8            `dd:"type:tinyint(4)"`
	CreateAt  *time.Time      `dd:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdateAt  *time.Time      `dd:"type:datetime;default:CURRENT_TIMESTAMP;extra:on update CURRENT_TIMESTAMP"`
	DeleteAt  *time.Time      `dd:"type:datetime"`
	AvgScore  decimal.Decimal `dd:"type:decimal(6,4)"`
	Hobby     string          `dd:"type:varchar(255)"`
}
