FROM loads/alpine:3.8

WORKDIR /usr/local/ws
ADD ./service_order /usr/local/ws
ADD config /usr/local/ws/config
RUN chmod +x /usr/local/ws/service_order
EXPOSE 8080
CMD ["./service_order"]