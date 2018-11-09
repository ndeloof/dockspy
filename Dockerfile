FROM alpine

RUN  apk --update --no-cache add strace socat

VOLUME /var/run
CMD socat -v UNIX-LISTEN:/var/run/docker.sock,fork UNIX:/var/run/real-docker.sock