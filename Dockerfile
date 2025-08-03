FROM golang:1.24-alpine AS builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags='-s -w' -o uwdog cmd/uwdog/main.go

FROM alpine:latest
COPY --from=builder /build/uwdog /uwdog/uwdog
WORKDIR /uwdog
ENTRYPOINT ["./uwdog"]