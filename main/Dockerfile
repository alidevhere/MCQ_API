# syntax=docker/dockerfile:1
FROM golang:1.17-alpine
WORKDIR /app
COPY go.mod /app/
COPY go.sum /app/
RUN go mod download
COPY *.go /app/
COPY data.csv /app/
RUN go build -o api
CMD ["./api"]