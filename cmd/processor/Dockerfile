FROM golang:1.23-alpine

# Установить рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum из корня проекта
COPY ../../go.mod ../../go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь код сервиса processor
COPY ./cmd/processor/ .

# Собираем бинарный файл processor
RUN go build -o processor

# Запуск сервиса processor
CMD ["./processor"]
