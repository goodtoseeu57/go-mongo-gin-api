FROM golang:1.20.4-buster

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils 2>&1 \
    && useradd -s /bin/bash -m serveruser

RUN apt-get -y install git iproute2 procps lsb-release

RUN chown -R serveruser:serveruser /go

USER serveruser

ENV DEBIAN_FRONTEND=dialog

EXPOSE 8000