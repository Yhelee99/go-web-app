package routes

import (
	"GoWebApp/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func SetUp() http.Handler {
	r := gin.New()
	r.Use(logger.GinLogger(zap.L()), logger.GinRecovery(zap.L(), true))

	r.GET("/", func(c *gin.Context) {
		zap.L().Debug("服务被访问！")
		c.JSON(http.StatusOK, gin.H{
			"msg": "Welcome to My First Web App!",
		})
	})
	return r
}
