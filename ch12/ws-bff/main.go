package main

import (
	"github.com/gin-gonic/gin"
	"go-grpc/ch12/ws-bff/handler"
	"go.elastic.co/apm/module/apmgin"
)

func main() {
	engine := initRouter()

	engine.Run(":9100")
}

func initRouter() *gin.Engine {
	engine := gin.Default()
	engine.Use(apmgin.Middleware(engine))

	engine.GET("/", handler.IndexApi)
	return engine
}
