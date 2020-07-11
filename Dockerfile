# build stage
FROM golang:alpine as builder
RUN apk add --no-cache git gcc libc-dev ca-certificates openssl
RUN go get github.com/Kong/go-pluginserver


RUN mkdir /go-plugins
COPY plugins/custom-auth.go /go-plugins/custom-auth.go
COPY plugins/api-key.go /go-plugins/api-key.go
RUN go build -buildmode plugin -o /go-plugins/custom-auth.so /go-plugins/custom-auth.go
RUN go build -buildmode plugin -o /go-plugins/api-key.so /go-plugins/api-key.go

# production stage
FROM kong:2.0.1-alpine
COPY --from=builder /go/bin/go-pluginserver /usr/local/bin/go-pluginserver
RUN mkdir /tmp/go-plugins
COPY --from=builder /go-plugins/custom-auth.so /tmp/go-plugins/custom-auth.so
COPY --from=builder /go-plugins/api-key.so /tmp/go-plugins/api-key.so
COPY config.yml /tmp/config.yml

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER root
RUN chmod -R 777 /tmp
USER kong