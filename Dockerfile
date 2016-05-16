FROM gliderlabs/alpine
MAINTAINER Kevin Stock <kevinstock@tantalic.com>

RUN apk-install ca-certificates
ADD ./twitter-service /usr/local/bin/

ENTRYPOINT /usr/local/bin/twitter-service
EXPOSE 3000
