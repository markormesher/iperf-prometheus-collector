FROM docker.io/golang:1.25.1@sha256:8305f5fa8ea63c7b5bc85bd223ccc62941f852318ebfbd22f53bbd0b358c07e1 as builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./cmd ./cmd

RUN go build -o ./build/main ./cmd/...

# ---

FROM docker.io/debian:13.1@sha256:833c135acfe9521d7a0035a296076f98c182c542a2b6b5a0fd7063d355d696be
WORKDIR /app

LABEL image.registry=ghcr.io
LABEL image.name=markormesher/iperf-prometheus-collector

RUN apt update \
  && apt install -y --no-install-recommends iperf3 \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/build/main /usr/local/bin/iperf-prometheus-collector

CMD ["/usr/local/bin/iperf-prometheus-collector"]
