FROM golang:latest

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

EXPOSE 8080

RUN apt-get update

CMD ["tail", "-f", "/dev/null"]