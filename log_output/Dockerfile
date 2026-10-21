FROM golang:1.21-alpine
WORKDIR /app
COPY main.go .
RUN go build -o log-output main.go
CMD ["./log-output"]