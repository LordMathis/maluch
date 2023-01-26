package main

import (
	"os"

	"github.com/LordMathis/maluch/web"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {

	server := web.SetupRouter()
	server.Run(":8080")
}
