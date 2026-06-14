package main

import (
	"github.com/MindlessMuse666/auth/internal/app"
	"github.com/MindlessMuse666/auth/internal/config"
	"github.com/MindlessMuse666/auth/internal/logging"
)

func main() {
	cfg := config.MustLoad()
	log := logging.New(cfg.Env)

	log.Info("Запуск приложения")
	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
	application.GRPCSrv.MustRun()
}
