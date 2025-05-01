FROM golang:1.23.8-alpine3.20 as build

WORKDIR /app

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:3.15 as runner

COPY --from=build /app/plugin-discovery /opt/plugin-discovery

WORKDIR /opt

EXPOSE 8080

CMD ["./plugin-discovery"]