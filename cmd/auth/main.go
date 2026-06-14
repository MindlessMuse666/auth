package main

import (
	"github.com/MindlessMuse666/auth/internal/config"
	"github.com/MindlessMuse666/auth/internal/logging"
)

func main() {
	cfg := config.MustLoad()
	log := logging.New(cfg.Env)

	log.Info(
		"Запуск приложения",
		// slog.Any("cfg", cfg),
	)

	// TODO: инициализировать аппку (пакет app)

	// TODO: запустить gRPC-сервер аппки
}
