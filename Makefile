# Переменные
GO = go
TEST_DIR_VALIDATEJSON = ./internal/command/validateJSON
TEST_FLAGS = -v

# Цель по умолчанию
.PHONY: all
all: validateJSON

# Запуск тестов validateJSON
.PHONY: validateJSON
validateJSON:
	$(GO) test $(TEST_DIR_VALIDATEJSON) $(TEST_FLAGS)

# Помощь
.PHONY: help
help:
	@echo "Доступные команды:"
	@echo "  make validateJSON  - Запустить тесты пакета validateJSON"
	@echo "  make help         - Показать эту справку"