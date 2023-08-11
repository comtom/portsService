# builder image
FROM golang:1.21.0-alpine3.18 as builder
RUN apk update && \
    apk add --no-cache git && \
    rm -rf /var/cache/apk/*
RUN mkdir /build
ADD . /build/
WORKDIR /build
ENV GOPROXY=direct
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ports-service .

# final image for distributing
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /build/ports-service .

ENTRYPOINT [ "./ports-service" ]
