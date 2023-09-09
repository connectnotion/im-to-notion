FROM alpine:latest as prod

WORKDIR /usr/local/im-to-notion

COPY ./bin/im-to-notion ./
COPY ./conf/config.yaml ./conf/config.yaml

ENTRYPOINT [ "/usr/local/im-to-notion/im-to-notion" ]
