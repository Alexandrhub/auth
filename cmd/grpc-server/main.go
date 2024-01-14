package main

import (
	"context"
	"log"

	"github.com/alexandrhub/auth/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %v", err.Error())
	}

	if err := a.Run(); err != nil {
		log.Fatalf("failed to run app: %v", err.Error())
	}
}
