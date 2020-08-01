package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	db()
	r := gin.Default()
	r.GET("/orders", getOrders)
	r.Run() // listen and serve on 0.0.0.0:8080
}
