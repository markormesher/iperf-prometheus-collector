FROM node:20.15.0-alpine@sha256:df01469346db2bf1cfc1f7261aeab86b2960efa840fe2bd46d83ff339f463665 AS builder

WORKDIR /iperf-prometheus-collector

COPY ./package.json ./yarn.lock ./
RUN yarn install

COPY ./tsconfig.json ./
COPY ./src ./src/
RUN yarn build

# ---

FROM node:20.15.0-alpine@sha256:df01469346db2bf1cfc1f7261aeab86b2960efa840fe2bd46d83ff339f463665

WORKDIR /iperf-prometheus-collector

RUN apk add --no-cache iperf3

COPY ./package.json ./yarn.lock ./
RUN yarn install --production

COPY --from=builder /iperf-prometheus-collector/build ./build/

EXPOSE 9030
CMD yarn start
