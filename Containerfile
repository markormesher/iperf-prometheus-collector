FROM docker.io/golang:1.25.0@sha256:5502b0e56fca23feba76dbc5387ba59c593c02ccc2f0f7355871ea9a0852cebe as builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./cmd ./cmd

RUN go build -o ./build/main ./cmd/...

# ---

FROM docker.io/debian:12.11@sha256:731dd1380d6a8d170a695dbeb17fe0eade0e1c29f654cf0a3a07f372191c3f4b
WORKDIR /app

LABEL image.registry=ghcr.io
LABEL image.name=markormesher/iperf-prometheus-collector

RUN apt update \
  && apt install -y --no-install-recommends iperf3 \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/build/main /usr/local/bin/iperf-prometheus-collector

CMD ["/usr/local/bin/iperf-prometheus-collector"]
