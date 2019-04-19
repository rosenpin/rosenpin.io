FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

ADD bin/rosenpin.app /bin/rosenpin.app
ADD configs /configs
ADD resources /resources

# Execute the binary on run
ENTRYPOINT rosenpin.app -c configs/production_config.yml 

# Expose port 80 for host
EXPOSE 8080
