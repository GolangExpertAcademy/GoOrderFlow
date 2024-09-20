FROM golang:1.23-alpine

# Установить рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum из корня проекта
COPY ../../go.mod ../../go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь код сервиса web
COPY ./cmd/web/ .

# Собираем бинарный файл web
RUN go build -o web

# Запуск сервиса web
CMD ["./web"]
