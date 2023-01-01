package main

import (
	"awesomeProject4/internal/api/darksky"
	"fmt"
	"log"
)

func main() {
	bytes, err := darksky.MakeRequest()
	if err != nil {
		log.Fatalf("error while receiving data from API: %s", err)
	}
	for v := range bytes {
		fmt.Println(string(v))
	}
}
