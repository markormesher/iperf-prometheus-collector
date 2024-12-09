FROM docker.io/golang:1.23.4@sha256:574185e5c6b9d09873f455a7c205ea0514bfd99738c5dc7750196403a44ed4b7 as builder
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
