package main

import (
	"advanced_programming/clients"
	"advanced_programming/config"
	"advanced_programming/constant"

	"github.com/gin-gonic/gin"
)

func main() {
	ReadConfig(constant.ProductFilePath)
	Init()
	r := gin.Default()
	Register(r)
	_ = r.Run()
}

func ReadConfig(file string) {
	config.ReadConfig(file)
}

func Init() {
	clients.InitMysql()
}

func Register(r *gin.Engine) {
	UserRegister(r)
	BlogRegister(r)
}
