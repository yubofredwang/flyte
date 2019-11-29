FROM ubuntu:bionic

# To build this, cd into the portal directory on this branch and run
#   docker build . -t docker.io/lyft/flyte:portal -f Dockerfile

WORKDIR /root
ENV VENV /root/venv
ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8
ENV PYTHONPATH /root

RUN : \
    && apt-get update && apt-get install -y python3.6 python3.6-venv python3.6-dev \
    && apt-get install -y make build-essential libssl-dev python3-venv python3-pip \
    && apt-get install -y curl wget dnsutils htop

RUN pip3 install awscli

WORKDIR /setup/redis
RUN wget http://download.redis.io/redis-stable.tar.gz \
    && tar xvzf redis-stable.tar.gz \
    && cd redis-stable \
    && make \
    && cp src/redis-cli /usr/local/bin/

RUN apt-get install -y postgresql-client

WORKDIR /root
