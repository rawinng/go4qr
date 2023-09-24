package utilities

import "encoding/json"

// implement CommonJson interface
func ToJson(o any) string {
	// marshal struct to json
	jsonBytes, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}

	return string(jsonBytes)
}

// Deserializes the JSON string into a struct
// , checkout Unmarshal function from https://golang.org/pkg/encoding/json/.
//
//	emptyObject must be a pointer
func FromJson(jsonString string, emptyObject any) {
	// unmarshal json to struct
	jsonBytes := []byte(jsonString)
	err := json.Unmarshal(jsonBytes, emptyObject)
	if err != nil {
		panic(err)
	}
}
