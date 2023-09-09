<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [im-to-notion](#im-to-notion)
  - [set config file](#set-config-file)
  - [build on local](#build-on-local)
  - [use release image](#use-release-image)
- [Acknowledgement](#acknowledgement)
- [Buy me a coffee](#buy-me-a-coffee)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# im-to-notion

send msg to notion through im.

## set config file

```yaml
log:
  level: info
  output: stdout

source:
  dingtalk:
    appKey: xxx
    appSecret: xxx

destination:
  notion:
    secret: xxx
    databaseId: xxx
```

## build on local

```shell
make build
```

```shell
# after modify ./conf/config.yaml
./bin/im-to-notion
```

## use release image

- use env

```shell
docker run -itd \
 -e LOG_LEVEL=${LOG_LEVEL} \
 -e LOG_OUTPUT=${LOG_OUTPUT} \
 -e DINGTALK_APP_KEY=${DINGTALK_APP_KEY} \
 -e DINGTALK_APP_SECRET=${DINGTALK_APP_SECRET} \
 -e NOTION_SECRET=${NOTION_SECRET} \
 -e NOTION_DATABASE_ID=${NOTION_DATABASE_ID} \
 --name im-to-notion ghcr.io/ronething/im-to-notion:0.0.1
```

- use config file

```shell
docker run -itd -v `pwd`/conf:/usr/local/im-to-notion/conf --name im-to-notion ghcr.io/ronething/im-to-notion:0.0.1 -c /usr/local/im-to-notion/conf/config.yaml
```

# Acknowledgement

- dingtalk
- notion
- opensource project in [go.mod](./go.mod)

# Buy me a coffee

<img src="https://github.com/ronething/ronething/blob/master/images/wechat.jpg?raw=true" alt="wechat" width="300px" />

