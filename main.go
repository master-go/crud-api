package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/master-go/crud-api/initializers"
	"gorm.io/gorm"
)

func main() {
	initializers.LoadEnvVeriables()
	var DB *gorm.DB = initializers.ConnectToDB()
	fmt.Println(DB)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
