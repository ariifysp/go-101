package custom

import (
	"encoding/json"
	"log"
)

func LogStruct(str string, v interface{}) {
	// Convert struct to JSON
	jsonData, err := json.Marshal(v)
	if err != nil {
		log.Fatalf("Error marshaling struct: %v", err)
	}

	log.Printf("%s: %s", str, string(jsonData))
}
