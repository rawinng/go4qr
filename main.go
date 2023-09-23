// Make main function says "hello, world!" to console
package main

import (
	"easy-redis/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	startServer()
}

func startServer() {
	r := gin.Default()
	// set up router for /qr
	api := r.Group("/api/v1")
	// setup route
	routes.SetupQrRouter(api)
	// run server
	r.Run()
}
