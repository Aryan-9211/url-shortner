FROM golang:1.21.9 as builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o url_shortener main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/url_shortener .

EXPOSE 8080

CMD ["./url_shortener"]
