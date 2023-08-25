package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	azureSegment := segmentAZURE()
	jsonBytes, err := json.Marshal(azureSegment)
	if err != nil {
		fmt.Println("[]")
	}
	fmt.Println(string(jsonBytes))
}
