# To build:
# $ docker run --rm -v $(pwd):/go/src/github.com/micahhausler/template -w /go/src/github.com/micahhausler/template golang:1.6  go build -v -a -tags netgo -installsuffix netgo -ldflags '-w'
# $ docker build -t micahhausler/template .
#
# To run:
# $ docker run micahhausler/template

FROM busybox

MAINTAINER Micah Hausler, <hausler.m@gmail.com>

COPY template /bin/template
RUN chmod 755 /bin/template

ENTRYPOINT ["/bin/template"]
