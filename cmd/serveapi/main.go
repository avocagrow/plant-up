package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/avocagrow/plant-up/internal/app"
)

var env = os.Getenv("RUN_MODE")

func main() {
	if env == "" {
		env = "dev"
	}
	app, err := app.NewApp(
		app.WithName(fmt.Sprintf("coffee-up-%s-apiserver", env)),
	)
	if err != nil {
		slog.Error("unable to create new api app")
		os.Exit(1)
	}

	app.Server.ListenAndServe()
}
