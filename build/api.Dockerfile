FROM alpine:latest
RUN apk add -U --no-cache ca-certificates

ENV CONFIG_PATH /root/.hello-cafe/config.yml
ENV PORT 8000

COPY configs/config.yml /root/.hello-cafe/config.yml
COPY bin/api /usr/bin/api

ENTRYPOINT ["/usr/bin/api"]