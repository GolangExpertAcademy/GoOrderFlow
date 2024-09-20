FROM golang:1.23-alpine

# Установить рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum из корня проекта
COPY ../../go.mod ../../go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь код сервиса notifier
COPY ./cmd/notifier/ .

# Собираем бинарный файл notifier
RUN go build -o notifier

# Запуск сервиса notifier
CMD ["./notifier"]
