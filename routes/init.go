package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/siskinc/srv-name-list/middlewares"
	"github.com/siskinc/srv-name-list/routes/list_item"
	"github.com/siskinc/srv-name-list/routes/list_item_hit"
	"github.com/siskinc/srv-name-list/routes/list_type"
	"github.com/siskinc/srv-name-list/routes/namespace"
)

func Init(router *gin.Engine) {
	router.Use(cors.Default())
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(middlewares.LoggerFormatter))
	nameListGroup := router.Group("/name-list")
	namespaceGroup := nameListGroup.Group("/namespace")
	listTypeGroup := nameListGroup.Group("/type")
	listItemGroup := nameListGroup.Group("/item")
	listItemHitGroup := nameListGroup.Group("/item-hit")
	// 命名空间
	{
		namespaceGroup.GET("/", namespace.QueryNamespace)
		namespaceGroup.POST("/", namespace.CreateNamespace)
		namespaceGroup.DELETE("/:id", namespace.DeleteNamespace)
		namespaceGroup.PATCH("/:id", namespace.UpdateNamespace)
	}
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
		listItemGroup.POST("/", list_item.CreateListItem)
		listItemGroup.DELETE("/:id", list_item.DeleteListItem)
		listItemGroup.PATCH("/:id", list_item.UpdateListItem)
	}
	// 命中名单项接口
	{
		listItemHitGroup.POST("/pre", list_item_hit.ItemHitPre) // 预计算，针对某一个名单
		listItemHitGroup.POST("/all", list_item_hit.ItemHitAll) // 针对所有名单
	}
}
