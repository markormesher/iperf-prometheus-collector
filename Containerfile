FROM docker.io/golang:1.24.5@sha256:14fd8a55e59a560704e5fc44970b301d00d344e45d6b914dda228e09f359a088 as builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./cmd ./cmd

RUN go build -o ./build/main ./cmd/...

# ---

FROM docker.io/debian:bookworm@sha256:d42b86d7e24d78a33edcf1ef4f65a20e34acb1e1abd53cabc3f7cdf769fc4082
WORKDIR /app

LABEL image.registry=ghcr.io
LABEL image.name=markormesher/iperf-prometheus-collector

RUN apt update \
  && apt install -y --no-install-recommends iperf3 \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/build/main /usr/local/bin/iperf-prometheus-collector

CMD ["/usr/local/bin/iperf-prometheus-collector"]
