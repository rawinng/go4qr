package businesses

import (
	"easy-redis/pkg/clients"
	"easy-redis/pkg/models"
	"time"

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
	created := time.Now().String()
	return models.Qr{
		QrCode:  qr_code,
		Created: created,
	}
}

func saveQr(qr models.Qr) {
	// save to redis
	redisClient := clients.GetRedisInstance()
	qrString := qr.ToJson()
	err := redisClient.SetWithExpire(qr.QrCode, qrString, 60)
	if err != nil {
		panic(err)
	}
}
