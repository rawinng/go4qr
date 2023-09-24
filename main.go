// Make main function says "hello, world!" to console
package main

import (
	"github.com/rawinng/go4qr/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	startServer()
}

func prepareServer() *gin.Engine {
	r := gin.Default()
	// set up router for /qr
	api := r.Group("/api/v1")
	// setup route
	routes.SetupQrRouter(api)
	return r
}

func startServer() {
	r := prepareServer()
	// run server
	r.Run()
}
