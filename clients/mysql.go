package clients

import (
	"advanced_programming/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	log.Printf("%v", config.APConfig)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.APConfig.Mysql.User,
		config.APConfig.Mysql.Password,
		config.APConfig.Mysql.Host,
		config.APConfig.Mysql.Port,
		config.APConfig.Mysql.DBName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Printf("连接数据库失败, error = " + err.Error())
	}
}
