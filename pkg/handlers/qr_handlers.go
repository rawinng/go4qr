package handlers

// import gin
import (
	"github.com/rawinng/go4qr/pkg/businesses"
	//gin
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetQrCode(c *gin.Context) {
	// output struct from businesses.GetNewQr() to Json
	c.JSON(http.StatusOK, businesses.GetNewQr())
}
