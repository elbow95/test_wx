package main

import (
	"fmt"
	"net/http"
	"wxcloudrun-golang/common"
	"wxcloudrun-golang/handler"
	"wxcloudrun-golang/middleware"

	"github.com/gin-gonic/gin"
)

func RouterRegister() {
	router := gin.Default()

	// no login router
	router.POST("/app/phone", common.HandlerWrapper(handler.GetPhone))
	router.POST("/app/user_info", common.HandlerWrapper(handler.GetUserInfo))
	router.POST("/app/user/register", common.HandlerWrapper(handler.Register))

	// 开启跨域
	router.Use(middleware.Cors())
	router.Use(middleware.Login())

	// 404
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	app := router.Group("/app")
	{
		user := app.Group("/user")
		{
			user.POST("/list", common.HandlerWrapper(handler.ListUser))
			user.POST("/add", common.HandlerWrapper(handler.AddUser))
			user.POST("/delete", common.HandlerWrapper(handler.DeleteUser))
			user.POST("/update", common.HandlerWrapper(handler.UpdateUser))
		}

		station := app.Group("/station")
		{
			station.POST("/list", common.HandlerWrapper(handler.ListStation))
			station.POST("/add", common.HandlerWrapper(handler.AddStation))
			station.POST("/delete", common.HandlerWrapper(handler.DeleteStation))
			station.POST("/update", common.HandlerWrapper(handler.UpdateStation))
		}

		oil := app.Group("/oil")
		{
			oil.POST("/list", common.HandlerWrapper(handler.ListOil))
			oil.POST("/add", common.HandlerWrapper(handler.AddOil))
			oil.POST("/delete", common.HandlerWrapper(handler.DeleteOil))
			oil.POST("/update", common.HandlerWrapper(handler.UpdateOil))
		}

		record := app.Group("/record")
		{
			record.POST("/list", common.HandlerWrapper(handler.ListRecord))
			record.POST("/add", common.HandlerWrapper(handler.AddRecord))
			record.POST("/delete", common.HandlerWrapper(handler.DeleteRecord))
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
