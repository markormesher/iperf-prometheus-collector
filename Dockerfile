FROM node:21.7.3-bookworm@sha256:4b232062fa976e3a966c49e9b6279efa56c8d207a67270868f51b3d155c4e33d AS builder
WORKDIR /app

COPY .yarn/ .yarn/
COPY package.json yarn.lock .yarnrc.yml ./
RUN yarn install

COPY ./src ./src
COPY ./tsconfig.json ./

RUN yarn build

# ---

FROM node:21.7.3-bookworm@sha256:4b232062fa976e3a966c49e9b6279efa56c8d207a67270868f51b3d155c4e33d
WORKDIR /app

RUN apt update && apt install -y --no-install-recommends iperf3

COPY .yarn/ .yarn/
COPY package.json yarn.lock .yarnrc.yml ./
RUN yarn workspaces focus --all --production

COPY --from=builder /app/build /app/build

EXPOSE 9030

USER nobody
CMD yarn start
