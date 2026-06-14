# Переменные
MAKEFILE	:= [MAKEFILE]
GO			:= go
PROTOC		:= protoc
PROTO_DIR	:= proto
GEN_DIR		:= gen/go
CONFIG_FILE	:= ./config/local.toml
MAIN_PKG	:= ./cmd/auth/main.go

# Флаги для protoc
PROTO_FLAGS := --go_out=$(GEN_DIR) --go_opt=paths=source_relative \
			   --go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative

.PHONY: run proto clean help
.DEFAULT_GOAL := run

run:
	@echo $(MAKEFILE) Запуск сервиса аутентификации...
	$(GO) run $(MAIN_PKG) --config=$(CONFIG_FILE)

proto: $(PROTO_DIR)/auth.proto
	@echo $(MAKEFILE) Генерация protobuf кода...
	$(PROTOC) -I $(PROTO_DIR) $< $(PROTO_FLAGS)

help:
	@echo Использование: make [цель]
	@echo Цели:
	@echo 	run	Запустить auth-сервис (по умолчанию)
	@echo 	proto	Сгенерировать код protobuf и gRPC
	@echo 	help	Показать эту справку
