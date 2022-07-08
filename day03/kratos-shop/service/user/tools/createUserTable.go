package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          int64      `gorm:"primarykey"`
	Mobile      string     `gorm:"index:idx_mobile;unique;type:varchar(11) comment '手机号码,唯一标识';not null"`
	Password    string     `gorm:"type:varchar(100) comment '密码需要加密';not null"`
	NickName    string     `gorm:"type:varchar(25) comment '用户昵称'"`
	Birthday    *time.Time `gorm:"type:datetime comment '用户生日'"`
	Gender      string     `gorm:"column:gender;default:male;type:varchar(16) comment '性别'"`
	Role        int        `gorm:"column:role;default:1;type:int comment '1 - 普通用户，2 - 管理员'"`
	CreatedAt   time.Time  `gorm:"column:create_time"`
	UpdatedAt   time.Time  `gorm:"column:update_time"`
	DeletedAt   gorm.DeletedAt
	isDeletedAt bool
}

var TableName = "users"

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/go_geek?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	m := db.Migrator()
	if err := m.CreateTable(&User{}); err != nil {
		panic(err)
	}

}
