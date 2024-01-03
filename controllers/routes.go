package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/golang-llm-server/config"
)

func BuildRoutes(
	conf *config.Config,
	routerEngine *gin.Engine,
) {

	controller := NewController(conf)

	r := routerEngine.Group("/v1")

	r.GET("/health", func(c *gin.Context) {
		FormatResponse(c, http.StatusOK, "OK", nil)
	})

	r.Group("/", controller.CustomerEnquiry())

}
