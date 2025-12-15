FROM docker.io/golang:1.25.5@sha256:36b4f45d2874905b9e8573b783292629bcb346d0a70d8d7150b6df545234818f as builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./cmd ./cmd

RUN go build -o ./build/main ./cmd/...

# ---

FROM docker.io/debian:13.2@sha256:0d01188e8dd0ac63bf155900fad49279131a876a1ea7fac917c62e87ccb2732d
WORKDIR /app

LABEL image.registry=ghcr.io
LABEL image.name=markormesher/iperf-prometheus-collector

RUN apt update \
  && apt install -y --no-install-recommends iperf3 \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/build/main /usr/local/bin/iperf-prometheus-collector

CMD ["/usr/local/bin/iperf-prometheus-collector"]
