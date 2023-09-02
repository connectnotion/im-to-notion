<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [im-to-notion](#im-to-notion)
  - [docker run](#docker-run)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# im-to-notion

send msg to notion through im.

## docker run

```shell
docker run -itd -v `pwd`/conf:/usr/local/im-to-notion/conf --name im-to-notion ronething/im-to-notion:dev -c /usr/local/im-to-notion/conf/config.yaml
```
