package main

import (
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
)

func RouterRegister() {
	router := gin.Default()

	// // 开启跨域
	// router.Use(middleware.Cors())

	// // 404
	// router.NoRoute(func(c *gin.Context) {
	// 	c.String(http.StatusNotFound, "404 not found")
	// })

	router.GET("/", service.IndexHandler)
	app := router.Group("/app")
	{
		app.GET("/count", service.CounterHandler)
	}
}
