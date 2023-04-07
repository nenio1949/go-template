package bootstrap

import (
	"fmt"
	"go-template/global"
	"log"
	"net/http"

	server "go-template/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRoute() *gin.Engine {
	router := gin.Default()

	// 使用jwt验证登录
	// router.Use(middleware.JWTAuth(global.AppGuardName))

	// 注册 api 分组路由
	apiGroup := router.Group("/api")
	router.GET("/WW_verify_RrIfOPmetHUilu6o.txt", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintln("RrIfOPmetHUilu6o"))
	})
	router.Static("/static", "./server/templates")
	router.Use(cors.Default())
	server.SetApiUserRoutes(apiGroup)
	server.SetApiRoleRoutes(apiGroup)
	server.SetApiDepartmentRoutes(apiGroup)
	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setupRoute()
	srv := &http.Server{
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

	str := fmt.Sprintf("服务启动成功 %s:%s", global.App.Config.App.AppUrl, global.App.Config.App.Port)
	println(str)
}
