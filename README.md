<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [im-to-notion](#im-to-notion)
  - [set config file](#set-config-file)
  - [build image on local](#build-image-on-local)
  - [use release image](#use-release-image)
- [Acknowledgement](#acknowledgement)

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

## build image on local

```shell
make build-docker-image
```

```shell
docker run -itd -v `pwd`/conf:/usr/local/im-to-notion/conf --name im-to-notion im-to-notion:dev -c /usr/local/im-to-notion/conf/config.yaml
```

## use release image

```shell
docker run -itd -v `pwd`/conf:/usr/local/im-to-notion/conf --name im-to-notion docker pull ghcr.io/ronething/im-to-notion:0.0.1 -c /usr/local/im-to-notion/conf/config.yaml
```

# Buy me a coffee

<img src="https://github.com/ronething/ronething/blob/master/images/wechat.jpg?raw=true" alt="wechat" style="zoom:33%;" />

# Acknowledgement

- dingtalk
- notion
- opensource project in [go.mod](./go.mod)
