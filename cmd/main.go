package main

import (
	"JARVIS-agent/router"
	"fmt"
)

func main() {
	const port = ":5500"
	server := router.SetRouter()

	err := server.Run(port)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
