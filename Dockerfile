FROM golang:1.15.5-alpine as builder

WORKDIR /var/www/go-app
COPY . ./
RUN go build

FROM alpine:3.12 AS app
ARG HOME="/var/www"
RUN addgroup -g 10001 www-data \
      && adduser -u 10001 -G www-data -h /home/www-data -D www-data
COPY --from=builder /var/www/go-app/go-app /usr/local/bin/go-app

USER www-data
WORKDIR ${HOME}
ENTRYPOINT ["go-app"]
