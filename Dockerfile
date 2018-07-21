# Multi-stage build setup (https://docs.docker.com/develop/develop-images/multistage-build/)

# Stage 1 (to create a "build" image, ~850MB)
FROM golang:latest AS builder
RUN go version

COPY . /go/src/gitlab.com/johntang/db-employee
WORKDIR /go/src/gitlab.com/johntang/db-employee/
RUN set -x && \
    go get github.com/golang/dep/cmd/dep && \
    dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o microsvc .

# Stage 2 (to create a downsized "container executable", ~7MB)

# If you need SSL certificates for HTTPS, replace `FROM SCRATCH` with:
#
#   FROM alpine:3.7
#   RUN apk --no-cache add ca-certificates
#
FROM scratch
WORKDIR /opt/
COPY --from=builder /go/src/gitlab.com/johntang/microsvc-test/microsvc .
EXPOSE 8080
ENTRYPOINT ["./microsvc"]
