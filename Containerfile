FROM docker.io/golang:1.22.3@sha256:f43c6f049f04cbbaeb28f0aad3eea15274a7d0a7899a617d0037aec48d7ab010 as builder
WORKDIR /app

COPY go.mod Makefile ./
COPY ./cmd ./cmd

RUN make build

# ---

FROM docker.io/golang:1.22.3-alpine
WORKDIR /app

RUN apk add --no-cache iperf3

COPY --from=builder /app/build /app/build

CMD ["/app/build/main"]
