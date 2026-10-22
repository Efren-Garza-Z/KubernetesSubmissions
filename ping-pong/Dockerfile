FROM golang:1.21-alpine
WORKDIR /app
COPY main.go .
RUN go build -o ping-pong main.go
CMD ["./ping-pong"]