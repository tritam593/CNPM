FROM mysql:latest as db

ENV MYSQL_ROOT_PASSWORD=1

# Sao chép các tệp SQL khởi tạo vào container (nếu có)
COPY ./init.sql /docker-entrypoint-initdb.d/


FROM golang:1.21.4-alpine as program

WORKDIR /app
COPY ./app ./

# RUN go mod download


RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build main.go

CMD ["./main"]