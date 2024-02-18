FROM ubuntu
MAINTAINER eiixy <eiixyw@gmail.com>

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates curl vim

WORKDIR /app
COPY bin/example /app/example
COPY configs/example.yaml /app/configs/example.yaml
RUN mkdir "/app/logs"

ENTRYPOINT ["/app/example"]
