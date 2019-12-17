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


# Virtual environment
RUN python3.6 -m venv ${VENV}
RUN ${VENV}/bin/pip install wheel

# Install Python dependencies
# Been having problems with other versions of pip reading from gemfury
RUN ${VENV}/bin/pip install pip==19.0.3
RUN ${VENV}/bin/pip install flytekit
RUN ${VENV}/bin/pip install ipython

# This is a script that enables a virtualenv, copy it to a better location
RUN cp ${VENV}/bin/flytekit_venv /opt/
RUN chmod a+x /opt/flytekit_venv

COPY .bash_aliases /root/
COPY .bashrc /root
COPY motd /etc
RUN chmod a+r /etc/motd

WORKDIR /root
