# builder image
FROM golang:1.18-alpine3.16 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ports-service .

# final image for distributing
FROM alpine:3.16
COPY --from=builder /build/ports-service .

ENTRYPOINT [ "./ports-service" ]
