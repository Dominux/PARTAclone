package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)


func CreateWork(c *gin.Context) {
	fmt.Printf("%+v\n", c)
	c.JSON(200, gin.H{
		"message": "Пошел на хуй!",
	})
}

