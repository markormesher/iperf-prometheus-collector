FROM docker.io/golang:1.23.2@sha256:cc637ce72c1db9586bd461cc5882df5a1c06232fd5dfe211d3b32f79c5a999fc as builder
WORKDIR /app

COPY go.mod Makefile ./
COPY ./cmd ./cmd

RUN make build

# ---

FROM docker.io/golang:1.23.2-bookworm
WORKDIR /app

RUN apt update && apt install -y --no-install-recommends iperf3 && rm -rf /var/lib/apt/lists

COPY --from=builder /app/build /app/build

CMD ["/app/build/main"]
