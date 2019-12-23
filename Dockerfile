FROM golang:alpine AS build-env

# Set go bin which doesn't appear to be set already.
ENV GOBIN /go/bin

RUN apk update && apk upgrade && \
apk add --no-cache bash git openssh

# build directories
ADD . /go/src/mcules/rspamd-influxdb
WORKDIR /go/src/mcules/rspamd-influxdb

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN dep ensure

# Build my app
RUN go build -o rspamdInfluxDB *.go

# final stage
FROM alpine
WORKDIR /app

MAINTAINER McUles <rspamd-influxdb@itstall.de>

ENV INTERVAL=3600 \
    INFLUXDB_DB="rspamd" \
    INFLUXDB_URL="http://influxdb:8086" \
    INFLUXDB_USER="DEFAULT" \
    INFLUXDB_PWD="DEFAULT" \
    RSPAMD_URL="https://RSPAMD-URL" \
    RSPAMD_PASSWORD="PASSWORD"

RUN apk add ca-certificates
COPY --from=build-env /go/src/mcules/rspamd-influxdb/rspamdInfluxDB /app/rspamdInfluxDB
ADD run.sh run.sh
CMD sh run.sh
