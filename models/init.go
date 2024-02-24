package models

import (
	"gorm.io/gorm"
	"log"
	"server/global"
)

var db = global.MysqlDB

func init() {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Println("数据表创建失败", err)
		return
	}
}

// User todo: 用户表
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(15);not null;comment:账号" json:"username"`
	Password string `gorm:"type:varchar(15);not null;comment:密码" json:"password"`
	Status   int    `gorm:"type:tinyint(1);check:status IN (0,1);not null;comment:状态" json:"status"`
}

//type Goods struct {
//	gorm.Model
//	VideoName string `gorm:"type:varchar(30);not null;comment:视频名称" json:"username"`
//	VideoName string `gorm:"type:varchar(30);not null;comment:视频名称" json:"username"`
//	VideoName string `gorm:"type:varchar(30);not null;comment:视频名称" json:"username"`
//}
