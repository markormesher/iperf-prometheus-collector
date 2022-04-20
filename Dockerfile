FROM node:16.14.2

RUN apt update \
  && apt install -y iperf3 \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /iperf-prometheus-collector

COPY ./package.json ./yarn.lock ./
RUN yarn install

COPY ./tsconfig.json ./
COPY ./src ./src/
RUN yarn build

EXPOSE 9030
CMD yarn start
