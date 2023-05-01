FROM golang:1.19-alpine as builder
RUN apk update
RUN apk add --no-cache git
RUN rm -rf /build mkdir /build
ADD . /build
WORKDIR /build

CMD ["./run-test.sh"]
