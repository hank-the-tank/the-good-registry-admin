package main

import (
	app "./app"
	"github.com/gin-gonic/gin"
)

func main() {
	Connection()
	app.Fetch()
	r := gin.Default()
	r.GET("/orders", app.GetOrders)
	r.Run() // listen and serve on 0.0.0.0:8080
}
