package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/golang-llm-server/config"
)

type Controllers struct {
	conf *config.Config
}

func NewController(conf *config.Config) *Controllers {
	return &Controllers{conf: conf}
}

func (c *Controllers) CustomerEnquiry() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
