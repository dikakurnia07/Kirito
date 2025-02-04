FROM golang:1.21-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o kirito .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/kirito .
EXPOSE 8080
CMD ["./kirito"]
