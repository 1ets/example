package services

import (
	"example/app/adapters/servers"

	"github.com/gin-gonic/gin"
)

// Global Middleware setup
func HttpMiddleware(middleware *gin.Engine) {
	middleware.Use(gin.Recovery())
}

func HttpRouter(route *gin.Engine) {
	route.POST("example", servers.HttpPostExample)
	route.GET("example-database", servers.HttpGetDatabaseExample)
	route.GET("example-redis", servers.HttpRedisExample)
	route.POST("example-grpc", servers.HttpGrpcExample)
	route.POST("example-rabbitmq-async", servers.HttpRabbitAsyncExample)
	route.POST("example-rabbitmq-sync", servers.HttpRabbitSyncExample)

	// SAGA EXAMPLE
	route.POST("example-saga-stateless", servers.HttpSagaStatelessExample)
	route.POST("example-saga-stateful", servers.HttpSagaStatefulExample)
}
