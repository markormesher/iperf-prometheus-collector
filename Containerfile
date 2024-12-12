FROM docker.io/golang:1.23.4@sha256:70031844b8c225351d0bb63e2c383f80db85d92ba894e3da7e13bcf80efa9a37 as builder
WORKDIR /app

COPY go.mod Makefile ./
COPY ./cmd ./cmd

RUN make build

# ---

FROM docker.io/golang:1.23.4-bookworm@sha256:ef30001eeadd12890c7737c26f3be5b3a8479ccdcdc553b999c84879875a27ce
WORKDIR /app

RUN apt update && apt install -y --no-install-recommends iperf3 && rm -rf /var/lib/apt/lists

COPY --from=builder /app/build /app/build

CMD ["/app/build/main"]
