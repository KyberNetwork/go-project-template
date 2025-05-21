## BUILDER
FROM golang:1.24 AS builder

WORKDIR /src

COPY . .

RUN go build -o app ./cmd/service-name


## DEPLOY
FROM debian:bookworm

RUN apt-get update && \
    apt install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /cmd

COPY --from=builder /src/app /cmd/app

COPY migrations migrations

ENTRYPOINT /cmd/app
