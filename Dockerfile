FROM golang:1.23-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum (если они есть) для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем весь проект в рабочую директорию
COPY . .

#EXPOSE "44044"

# Устанавливаем air для горячей перезагрузки при разработке
RUN go install github.com/air-verse/air@latest

# Устанавливаем delve для отладки
#RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Команда для запуска приложения с использованием air для горячей перезагрузки
CMD ["air"]