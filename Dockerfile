FROM alpine

RUN  apk --update --no-cache add socat

VOLUME /var/run
CMD socat -v UNIX-LISTEN:/var/run/docker.sock,fork UNIX:/real-docker.sock