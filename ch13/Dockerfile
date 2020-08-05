FROM golang:1.13.5 AS builder

WORKDIR /app
COPY ./generate_config.go .
COPY ./docker-entrypoint.sh .
RUN CGO_ENABLED=0 GOOS=linux go build -o generate_config --ldflags "-extldflags -static" generate_config.go


FROM envoyproxy/envoy-alpine:v1.14.1

WORKDIR /etc
COPY --from=builder /app/generate_config .
COPY --from=builder /app/docker-entrypoint.sh .
RUN chmod +x /etc/docker-entrypoint.sh && chmod +x /etc/generate_config
ENTRYPOINT ["/etc/docker-entrypoint.sh"]
