package main

import infra "github.com/apirator/apirator-backend/internal/logger"

func main() {
	defer infra.Logger.Sync()
	app, err := Application()
	if err != nil {
		infra.Logger.Fatalw("application startup error", "error", err)
		panic("application startup error")
	}
	app.Run()
}
