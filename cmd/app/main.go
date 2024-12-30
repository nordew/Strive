package main

import (
	"github.com/joho/godotenv"
	"github.com/nordew/Strive/internal/app"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load env variables: %v", err)
	}
}

func main() {
	app.MustRun()
}
