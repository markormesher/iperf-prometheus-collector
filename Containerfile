FROM docker.io/golang:1.27.3@sha256:b76020929b709ddb6202afc05410a0faefae20ae05e27ad548b3870092688a69 as builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./cmd ./cmd

RUN go build -o ./build/main ./cmd/...

# ---

FROM docker.io/debian:13.7@sha256:2dc419c6d7fe3dad0a2f9fd469c34cc90b0f9ba642d8fde837fe3d7d34a0e7f1
WORKDIR /app

RUN apt update \
  && apt install -y --no-install-recommends iperf3 \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/build/main /usr/local/bin/iperf-prometheus-collector

CMD ["/usr/local/bin/iperf-prometheus-collector"]

LABEL image.name=markormesher/iperf-prometheus-collector
LABEL image.registry=ghcr.io
LABEL org.opencontainers.image.description=""
LABEL org.opencontainers.image.documentation=""
LABEL org.opencontainers.image.title="iperf-prometheus-collector"
LABEL org.opencontainers.image.url=""
LABEL org.opencontainers.image.vendor=""
LABEL org.opencontainers.image.version=""
