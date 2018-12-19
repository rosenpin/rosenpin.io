FROM golang


RUN go get -u gitlab.com/rosenpin/rosenpin.io/...
WORKDIR /go/src/gitlab.com/rosenpin/rosenpin.io
RUN go install gitlab.com/rosenpin/rosenpin.io/cmd/rosenpin/
ENTRYPOINT /go/bin/rosenpin.io -c /go/src/gitlab.com/rosenpin/rosenpin.io/production_config.yml

EXPOSE 80
EXPOSE 8080
#RUN /rosenpin.io/app -c rosenpin.io/production_config.yml
#CMD ["./app", "-c config.yml"]
