FROM alpine

RUN  apk --update --no-cache add socat

VOLUME /var/run
EXPOSE 2375
CMD socat -v TCP4-LISTEN:2375,fork,reuseaddr UNIX-CONNECT:/var/run/docker.sock
