package main

import infra "github.com/apirator/apirator-backend/internal/logger"

func main() {
	defer infra.Logger.Sync()
	app, _ := Application()
	app.Run()
}
