FROM golang

COPY spy.go spy.go

VOLUME /var/run
CMD go run spy.go
