FROM docker.io/golang:1.23.2@sha256:ad5c126b5cf501a8caef751a243bb717ec204ab1aa56dc41dc11be089fafcb4f as builder
WORKDIR /app

COPY go.mod Makefile ./
COPY ./cmd ./cmd

RUN make build

# ---

FROM docker.io/golang:1.23.2-alpine@sha256:9dd2625a1ff2859b8d8b01d8f7822c0f528942fe56cfe7a1e7c38d3b8d72d679
WORKDIR /app

RUN apk add --no-cache iperf3

COPY --from=builder /app/build /app/build

CMD ["/app/build/main"]
