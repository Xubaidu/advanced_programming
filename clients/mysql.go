package clients

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitMysql() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)")
	if err!= nil{
		panic(err)
	}
	defer db.Close()
}
