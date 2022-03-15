package domain

import (
	"time"
)

//dd:table
type InvalidToken struct {
	Id       int        `dd:"pk;auto;type:int(11)"`
	Token    string     `dd:"type:varchar(120);extra:comment 'revoked token';unique"`
	ExpireAt *time.Time `dd:"type:datetime;extra:comment 'expire time'"`
	CreateAt *time.Time `dd:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdateAt *time.Time `dd:"type:datetime;default:CURRENT_TIMESTAMP;extra:on update CURRENT_TIMESTAMP"`
	DeleteAt *time.Time `dd:"type:datetime"`
}
