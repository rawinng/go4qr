package routes

import (
	"github.com/rawinng/go4qr/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func SetupQrRouter(routerGroup *gin.RouterGroup) {
	// set up router for /qr
	routerGroup.GET("/qr", handlers.GetQrCode)
}
