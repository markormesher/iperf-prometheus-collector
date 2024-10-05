FROM node:20.18.0-alpine@sha256:c13b26e7e602ef2f1074aef304ce6e9b7dd284c419b35d89fcf3cc8e44a8def9 AS builder

WORKDIR /iperf-prometheus-collector

COPY ./package.json ./yarn.lock ./
RUN yarn install

COPY ./tsconfig.json ./
COPY ./src ./src/
RUN yarn build

# ---

FROM node:20.18.0-alpine@sha256:c13b26e7e602ef2f1074aef304ce6e9b7dd284c419b35d89fcf3cc8e44a8def9

WORKDIR /iperf-prometheus-collector

RUN apk add --no-cache iperf3

COPY ./package.json ./yarn.lock ./
RUN yarn install --production

COPY --from=builder /iperf-prometheus-collector/build ./build/

EXPOSE 9030
CMD yarn start
