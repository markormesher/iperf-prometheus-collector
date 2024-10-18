FROM golang:1.22.3@sha256:f43c6f049f04cbbaeb28f0aad3eea15274a7d0a7899a617d0037aec48d7ab010 as builder
WORKDIR /app

COPY go.mod Makefile ./
COPY ./cmd ./cmd

RUN make build

# ---

FROM golang:1.22.3-bookworm
WORKDIR /app

RUN apt update && apt install -y --no-install-recommends iperf3 && rm -rf /var/lib/apt/lists

COPY --from=builder /app/build /app/build

CMD ["/app/build/main"]
