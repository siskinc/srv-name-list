package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/siskinc/srv-name-list/middlewares"
	"github.com/siskinc/srv-name-list/routes/list_item"
	"github.com/siskinc/srv-name-list/routes/list_type"
)

func Init(router *gin.Engine) {
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(middlewares.LoggerFormatter))
	nameListGroup := router.Group("/name-list")
	listTypeGroup := nameListGroup.Group("/type")
	listItemGroup := nameListGroup.Group("/item")
	// 名单类型接口
	{
		listTypeGroup.GET("/", list_type.QueryListType)
		listTypeGroup.POST("/", list_type.CreateListType)
		listTypeGroup.DELETE("/:id", list_type.DeleteListType)
		listTypeGroup.PATCH("/:id", list_type.UpdateListType)
	}
	// 名单项接口
	{
		listItemGroup.GET("/", list_item.QueryListItem)
	}
}
