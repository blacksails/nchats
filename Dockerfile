FROM node:11 as app-builder

RUN mkdir /app
WORKDIR /app

COPY app/package.json app/yarn.lock ./
RUN yarn install 

COPY app .

RUN yarn run build


FROM golang:1.11-alpine as server-builder

RUN apk add --no-cache git build-base && \
  mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build ./cmd/nchats


FROM alpine

RUN mkdir /app

COPY --from=app-builder /app/dist /app/dist
COPY --from=server-builder /app/nchats /nchats

CMD ["/nchats"]
