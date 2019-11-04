FROM        alpine:3.10.3

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

COPY ghcs /bin/ghcs

ENTRYPOINT [ "/bin/ghcs" ]