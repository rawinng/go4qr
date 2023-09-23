package models

import "encoding/json"

// generate Json struct for this sample { qr: "qr_code" }
type Qr struct {
	QrCode  string `json:"qr_code"`
	Created string `json:"created"`
}

// implement CommonJson interface
func (qr *Qr) ToJson() string {
	// marshal struct to json
	jsonBytes, err := json.Marshal(qr)
	if err != nil {
		panic(err)
	}

	return string(jsonBytes)
}

func (qr Qr) FromJson(jsonString string) {
	// unmarshal json to struct
	jsonBytes := []byte(jsonString)
	err := json.Unmarshal(jsonBytes, &qr)
	if err != nil {
		panic(err)
	}
}
