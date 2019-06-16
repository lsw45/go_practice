package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// func main() {
// 	router := gin.Default()
// 	router.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "Hello World")
// 	})
// 	router.Run(":8000")
// }

func main() {
	router := gin.Default()
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	router.Run(":8000")
}
