FROM alpine:3.13

LABEL name="dockerhub_ratelimit_exporter" version="1.0.1" \
      description="A exporter for prometheus to check the pull limit of the DockerHub" \
      url="https://github.com/dohq/dockerhub_ratelimit_exporter"

RUN apk add --no-cache bash

COPY ./dockerhub_ratelimit_exporter /usr/local/bin
COPY ./entrypoint.sh /usr/local/bin

ENTRYPOINT ["entrypoint.sh"]
CMD ["dockerhub_ratelimit_exporter"]