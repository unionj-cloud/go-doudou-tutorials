package domain

import (
	"time"
)

//dd:table
type User struct {
	Id       int        `dd:"pk;auto;type:int(11)"`
	Username string     `dd:"type:varchar(45);extra:comment 'username';unique"`
	Password string     `dd:"type:varchar(60);extra:comment 'password'"`
	Name     string     `dd:"type:varchar(45);extra:comment 'real name'"`
	Phone    string     `dd:"type:varchar(45);extra:comment 'phone number'"`
	Dept     string     `dd:"type:varchar(45);extra:comment 'department'"`
	Avatar   string     `dd:"type:varchar(255);extra:comment 'user avatar'"`
	CreateAt *time.Time `dd:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdateAt *time.Time `dd:"type:datetime;default:CURRENT_TIMESTAMP;extra:on update CURRENT_TIMESTAMP"`
	DeleteAt *time.Time `dd:"type:datetime"`
}
