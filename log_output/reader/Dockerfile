FROM golang:1.21-alpine
WORKDIR /app
COPY main.go .
RUN go build -o reader main.go
CMD ["./reader"]