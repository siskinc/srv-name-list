package routes

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/global"
	"github.com/siskinc/srv-name-list/middlewares"
	"github.com/siskinc/srv-name-list/routes/list_item"
	"github.com/siskinc/srv-name-list/routes/list_item_hit"
	"github.com/siskinc/srv-name-list/routes/list_type"
	"github.com/siskinc/srv-name-list/routes/namespace"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// swagger
	{
		url := ginSwagger.URL("http://" + global.Config.SwaggerHost + "/swagger/doc.json") // The url pointing to API definition
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	// frontend
	{
		staticBox := packr.New("static", "../frontend/static")
		router.StaticFS("/static", staticBox)
		frontendBox := packr.New("favicon.ico", "../frontend")
		// router.StaticFS("/favicon.ico", frontendBox)
		t := template.New("tmp")
		var err error
		t, err = loadTemplate(t)
		if err != nil {
			logrus.Fatalf("load template have an err: %v", err)
		}
		router.SetHTMLTemplate(t)
		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})
		router.GET("/favicon.ico", gin.WrapH(http.FileServer(frontendBox)))
		// frontBox := packr.New("frontend", "../frontend")
		// router.StaticFS("/", frontBox)
	}
}

func loadTemplate(t *template.Template) (*template.Template, error) {
	box := packr.New("tmp", "../frontend")
	for _, file := range box.List() {
		if !strings.HasSuffix(file, ".html") {
			continue
		}
		h, err := box.FindString(file)
		if err != nil {
			return nil, err
		}
		//拼接方式，组装模板  admin/index.html 这种，方便调用
		t, err = t.New(file).Parse(h)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
