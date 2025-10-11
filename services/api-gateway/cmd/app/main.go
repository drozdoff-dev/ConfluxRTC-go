package main

import (
	"api-gateway/internal/app"
)

func main() {
	// Application entry point
	application := app.New()
	application.Run()
}
