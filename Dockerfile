FROM golang:1.17-alpine3.15 as build

WORKDIR /src

ADD . .

RUN apk add --no-cache make git && make build

FROM alpine:3.15

LABEL maintainer="wangrenzheng <renzheng.wang@ucloud.cn>"

WORKDIR /lxcfs

COPY --from=build /src/build/lxcfs-admission-webhook /lxcfs/lxcfs-admission-webhook

EXPOSE 8443

ENTRYPOINT ["/lxcfs/lxcfs-admission-webhook"]
