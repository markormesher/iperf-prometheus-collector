FROM node:16.14.2-alpine AS builder

WORKDIR /iperf-prometheus-collector

COPY ./package.json ./yarn.lock ./
RUN yarn install

COPY ./tsconfig.json ./
COPY ./src ./src/
RUN yarn build

# ---

FROM node:16.14.2-alpine

WORKDIR /iperf-prometheus-collector

RUN apk add --no-cache iperf3

COPY ./package.json ./yarn.lock ./
RUN yarn install --production

COPY --from=builder /iperf-prometheus-collector/build ./build/

EXPOSE 9030
CMD yarn start
