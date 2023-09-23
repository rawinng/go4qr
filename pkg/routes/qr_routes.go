package routes

import (
	"easy-redis/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func SetupQrRouter(routerGroup *gin.RouterGroup) {
	// set up router for /qr
	routerGroup.GET("/qr", handlers.GetQrCode)
}
