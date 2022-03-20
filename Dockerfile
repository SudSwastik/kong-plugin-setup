FROM golang:1.16-alpine as builder

RUN mkdir /go-plugins
COPY go-setheader.go /go-plugins/

RUN cd /go-plugins && \
    go mod init go-setheader && \
    go get -d -v github.com/Kong/go-pdk && \
    go get -d -v github.com/Kong/go-pdk/server && \
    go build .
    
FROM kong:2.8.0-alpine

USER root
RUN mkdir /go-plugins
COPY --from=builder /go-plugins/go-setheader /usr/local/bin/
USER kong