FROM golang


ADD . /go/src/gitlab.com/rosenpin/rosenpin.io
WORKDIR /go/src/gitlab.com/rosenpin/rosenpin.io
RUN go get ./...
RUN go install gitlab.com/rosenpin/rosenpin.io/cmd/rosenpin/
ENTRYPOINT /go/bin/rosenpin -c /go/src/gitlab.com/rosenpin/rosenpin.io/production_config.yml

EXPOSE 80
EXPOSE 8080
#RUN /rosenpin.io/app -c rosenpin.io/production_config.yml
#CMD ["./app", "-c config.yml"]
