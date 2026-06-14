package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/MindlessMuse666/auth/internal/app/grpc_app"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: инит хранилище (storage)

	// TODO: инит сервиса auth (auth)

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
