package main

import (
	"fmt"
	"net/http"
	"wxcloudrun-golang/handler"
	"wxcloudrun-golang/middleware"

	"github.com/gin-gonic/gin"
)

func RouterRegister() {
	router := gin.Default()

	// 开启跨域
	router.Use(middleware.Cors())

	// 404
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	app := router.Group("/app")
	{
		user := app.Group("/user")
		{
			user.GET("/list", handler.ListUser)
			user.POST("/add", handler.AddUser)
			user.POST("/delete", handler.DeleteUser)
			user.POST("/update", handler.UpdateUser)
		}

		station := app.Group("/station")
		{
			station.GET("/list", handler.ListStation)
			station.POST("/add", handler.AddStation)
			station.POST("/delete", handler.DeleteStation)
			station.POST("/update", handler.UpdateStation)
		}

		oil := app.Group("/oil")
		{
			oil.GET("/list", handler.ListOil)
			oil.POST("/add", handler.AddOil)
			oil.POST("/delete", handler.DeleteOil)
			oil.POST("/update", handler.UpdateOil)
		}

		record := app.Group("/record")
		{
			record.GET("/list", handler.ListRecord)
			record.POST("/add", handler.AddRecord)
			record.POST("/delete", handler.DeleteRecord)
		}
	}

	pc := router.Group("/pc")
	{
		record := pc.Group("/record")
		{
			record.GET("/list")
		}

		data := pc.Group("/data")
		{
			data.GET("/pie")
		}
	}

	_ = router.Run(fmt.Sprintf(":%d", 80))
}
