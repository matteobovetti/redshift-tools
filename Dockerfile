FROM alpine:latest AS deps

RUN apk --no-cache add ca-certificates

# base go image
FROM golang:1.27-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o redshift-tool ./cmd/redshift-tool/main.go

RUN chmod +x /app/redshift-tool

# build a tiny docker image
FROM alpine:latest

COPY --from=deps /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

RUN mkdir /app

COPY --from=builder /app/redshift-tool /app

ENTRYPOINT [ "/app/redshift-tool" ]
