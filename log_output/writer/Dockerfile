FROM golang:1.21-alpine
WORKDIR /app
COPY main.go .
RUN go build -o writer main.go
CMD ["./writer"]