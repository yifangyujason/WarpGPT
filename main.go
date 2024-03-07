package main

import (
	"WarpGPT/pkg/logger"
	"WarpGPT/pkg/db"
	"WarpGPT/pkg/env"
	"WarpGPT/pkg/funcaptcha"
	"WarpGPT/pkg/plugins"
	"WarpGPT/pkg/plugins/api/arkosetoken"
	"WarpGPT/pkg/plugins/api/backendapi"
	"WarpGPT/pkg/plugins/api/officialapi"
	"WarpGPT/pkg/plugins/api/publicapi"
	"WarpGPT/pkg/plugins/api/rapi"
	"WarpGPT/pkg/plugins/api/session"
	"WarpGPT/pkg/plugins/api/unofficialapi"
	"WarpGPT/pkg/plugins/service/proxypool"
	"github.com/bogdanfinn/fhttp"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("AuthKey")
		if apiKey != env.E.AuthKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Next()
	}
}
func main() {
	var router = gin.Default()
	if env.E.Verify {
		router.Use(AuthMiddleware())
	}
	router.Use(CORSMiddleware())
	component := &plugins.Component{
		Engine: router,
		Logger: logger.Log,
		Env:    &env.E,
		Auth:   funcaptcha.GetOpenAIArkoseToken,
		Db: db.DB{
			GetRedisClient: db.GetRedisClient,
		},
	}
	var plugin_list []plugins.Plugin
	plugin_list = append(
		plugin_list,
		&arkosetoken.ArkoseTokenInstance,
		&session.SessionTokenInstance,
		&backendapi.BackendProcessInstance,
		&officialapi.OfficialApiProcessInstance,
		&unofficialapi.UnofficialApiProcessInstance,
		&publicapi.PublicApiProcessInstance,
		&rapi.ApiProcessInstance,
		&proxypool.ProxyPoolInstance,
	)
	for _, plugin := range plugin_list {
		plugin.Run(component)
	}
	router.Run(env.E.Host + ":" + strconv.Itoa(env.E.Port))
}
