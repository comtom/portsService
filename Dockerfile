# builder image
FROM golang:1.13-alpine3.11 as builder
RUN mkdir /build
ADD *.go /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ports-service .

# final image for distributing
FROM alpine:3.11.3
COPY --from=builder /build/ports-service .

ENTRYPOINT [ "./ports-service" ]
