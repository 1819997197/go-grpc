package handler

import (
	"github.com/gin-gonic/gin"
	"go-grpc/ch12/ws-bff/ext_api"
	"net/http"
)

func IndexApi(c *gin.Context) {
	result, err := ext_api.SayHello(c.Request.Context())
	if err != nil {
		c.String(http.StatusOK, "it fail!")
		return
	}

	c.String(http.StatusOK, "it ok!"+result)
}
