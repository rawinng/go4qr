package models

// generate Json struct for this sample { qr: "qr_code" }
type Qr struct {
	QrCode    string `json:"qr_code"`
	CreatedAt string `json:"created"`
}
