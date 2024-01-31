FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o Web-App-Messenger .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/Web-App-Messenger .

CMD ["./Web-App-Messenger"]