package main

import (
	"advanced_programming/clients"

	"github.com/gin-gonic/gin"
)

func main() {
	Init()
	r := gin.Default()
	Register(r)
	r.Run()
}

func Init() {
	clients.InitMysql()
}

func Register(r *gin.Engine) {
	BlogRegister(r)
}