FROM golang:1.15.2-buster

RUN apt-get update -qq && \
    apt-get install -y mariadb-client vim

ENV GO111MODULE=on

COPY . /go/src/app/
WORKDIR /go/src/app

CMD ["./scripts/start_app.sh"]