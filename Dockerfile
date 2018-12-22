FROM golang

# Copy the source code including configuration
ADD . /go/src/gitlab.com/rosenpin/rosenpin.io
# Set the working directory
WORKDIR /go/src/gitlab.com/rosenpin/rosenpin.io
# Install deps
RUN go get ./...
# Build the binary
RUN go install gitlab.com/rosenpin/rosenpin.io/cmd/rosenpin/

# Execute the binary on run
ENTRYPOINT /go/bin/rosenpin -c /go/src/gitlab.com/rosenpin/rosenpin.io/configs/production_config.yml

# Expose port 80 for host
EXPOSE 80
