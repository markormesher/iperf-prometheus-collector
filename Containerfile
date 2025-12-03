FROM docker.io/golang:1.25.5@sha256:20b91eda7a9627c127c0225b0d4e8ec927b476fa4130c6760928b849d769c149 as builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./cmd ./cmd

RUN go build -o ./build/main ./cmd/...

# ---

FROM docker.io/debian:13.2@sha256:8f6a88feef3ed01a300dafb87f208977f39dccda1fd120e878129463f7fa3b8f
WORKDIR /app

LABEL image.registry=ghcr.io
LABEL image.name=markormesher/iperf-prometheus-collector

RUN apt update \
  && apt install -y --no-install-recommends iperf3 \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/build/main /usr/local/bin/iperf-prometheus-collector

CMD ["/usr/local/bin/iperf-prometheus-collector"]
