FROM library/node:14.19.0

MAINTAINER karocxing@163.com

RUN npm update \
  && npm upgrade \
  && npm install -g apidoc