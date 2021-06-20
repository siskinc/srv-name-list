package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/siskinc/srv-name-list/global"
	"github.com/siskinc/srv-name-list/routes"
)

func main() {
	router := gin.Default()
	routes.Init(router)
	gin.Logger()
	router.Run(":8000")
}
