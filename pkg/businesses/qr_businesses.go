package businesses

import (
	"strings"
	"time"

	"github.com/rawinng/go4qr/pkg/clients"
	"github.com/rawinng/go4qr/pkg/models"
	"github.com/rawinng/go4qr/pkg/utilities"

	"github.com/google/uuid"
)

func GetNewQr() *models.Qr {
	// generate new QR
	qr := generateQr()
	// save to redis
	saveQr(qr)
	return &qr
}

func generateQr() models.Qr {
	qr_code := uuid.New().String()
	// clear '-' from qr_code
	qr_code = strings.Replace(qr_code, "-", "", -1)
	created_time := time.Now()
	return models.Qr{
		QrCode:    qr_code,
		CreatedAt: utilities.FormatDefaultTime(created_time),
	}
}

// saveQr save qr to redis
func saveQr(qr models.Qr) {
	// save to redis
	redisClient := clients.GetRedisInstance()
	qrString := utilities.ToJson(qr)
	err := redisClient.SetWithExpire(qr.QrCode, qrString, 60)
	if err != nil {
		panic(err)
	}
}
