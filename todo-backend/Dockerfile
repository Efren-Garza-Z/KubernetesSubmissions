FROM golang:1.21-alpine
WORKDIR /app
COPY main.go .
RUN go build -o todo-backend main.go
CMD ["./todo-backend"]