package domain

import (
	"time"
)

//dd:table
type WordCloudTask struct {
	Id       int        `dd:"pk;auto;type:int(11)"`
	SrcUrl   string     `dd:"type:varchar(255);extra:comment 'source file url 原文件链接'"`
	ImgUrl   string     `dd:"type:varchar(255);extra:comment 'word cloud image url 词云图链接'"`
	Status   int8       `dd:"extra:comment 'task status 任务状态, 0 waiting 未开始, 1 processing 进行中, 2 success 成功, 3 fail 失败'"`
	Error    string     `dd:"type:varchar(655);extra:comment 'error message 错误信息'"`
	UserId   int        `dd:"type:int(11);fk:t_user,id,fk_user,ON DELETE CASCADE ON UPDATE NO ACTION"`
	CreateAt *time.Time `dd:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdateAt *time.Time `dd:"type:datetime;default:CURRENT_TIMESTAMP;extra:on update CURRENT_TIMESTAMP"`
	DeleteAt *time.Time `dd:"type:datetime"`
}
