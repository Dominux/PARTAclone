package main

import (
	"github.com/gin-gonic/gin"

	"github.com/SPARTAclone/SPARTA/internal/homeworks/controllers"
)


func main() {
	r := gin.Default()
	r.GET("/", controllers.CreateWork)
	r.Run()
}
