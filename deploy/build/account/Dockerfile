FROM ubuntu
MAINTAINER eiixy <eiixyw@gmail.com>

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates curl vim

WORKDIR /app
COPY bin/account /app/account
COPY configs/account.yaml /app/configs/account.yaml
RUN mkdir "/app/logs"

ENTRYPOINT ["/app/account"]
