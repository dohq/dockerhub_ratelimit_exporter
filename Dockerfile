# Builder Stage
FROM golang:1.20-alpine as builder

WORKDIR /go/src/app
COPY . /go/src/app

RUN mkdir /go/out && go build -o /go/out/dockerhub_ratelimit_exporter


# Final Image
FROM alpine:3.18

LABEL name="dockerhub_ratelimit_exporter" version="1.0.2" \
      description="A exporter for prometheus to check the pull limit of the DockerHub" \
      url="https://github.com/dohq/dockerhub_ratelimit_exporter"

RUN apk add --no-cache bash

COPY --from=builder /go/out/dockerhub_ratelimit_exporter /usr/local/bin
COPY ./entrypoint.sh /usr/local/bin

ENTRYPOINT ["entrypoint.sh"]
CMD ["dockerhub_ratelimit_exporter"]