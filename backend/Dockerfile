# Используем базовый образ Golang
FROM golang:1.22 AS builder

# Устанавливаем переменную окружения для корректной работы Go модулей
ENV GO111MODULE=on

# Устанавливаем рабочую директорию
WORKDIR /augventure-app

# Копируем go.mod и go.sum и устанавливаем зависимости
COPY go.mod .
COPY go.sum .
RUN go mod download

# Копируем остальные файлы проекта
COPY . .

# Собираем бинарный файл
RUN go build -o augventure-app cmd/main.go

# Используем легковесный образ для запуска бинарного файла
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем бинарный файл из предыдущего этапа
COPY --from=builder /augventure-app/augventure-app .
COPY --from=builder /augventure-app/schema ./schema

# Открываем порт для доступа к бэкенду
EXPOSE 8000

# Запускаем бэкенд
CMD ["./augventure-app"]