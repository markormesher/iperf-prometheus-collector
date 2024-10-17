FROM node:21.7.3-alpine AS builder
WORKDIR /app

COPY .yarn/ .yarn/
COPY package.json yarn.lock .yarnrc.yml ./
RUN yarn install

COPY ./src ./src
COPY ./tsconfig.json ./

RUN yarn build

# ---

FROM node:21.7.3-alpine
WORKDIR /app

RUN apk add --no-cache iperf3

COPY .yarn/ .yarn/
COPY package.json yarn.lock .yarnrc.yml ./
RUN yarn workspaces focus --all --production

COPY --from=builder /app/build /app/build

EXPOSE 9030

USER nobody
CMD yarn start
