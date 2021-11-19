FROM golang:1.17.3-alpine AS builder
RUN mkdir /build
ADD go.mod go.sum main.go constants.go /build/
WORKDIR /build
RUN go build

FROM alpine:3.14
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/twitch-command-api /app/
WORKDIR /app
CMD ["./twitch-command-api"]