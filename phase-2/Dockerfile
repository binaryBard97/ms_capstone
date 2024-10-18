FROM golang:bookworm AS builder

RUN apt-get update && \
    apt-get install -y \
    git \
    build-essential \
    pkg-config \
    libprotobuf-dev \
    protobuf-compiler \
    libnl-3-dev \
    libcap-dev \
    python3 \
    python3-protobuf \
    libprotobuf-c-dev \
    libprotobuf-c1 \
    libnet-dev \
    libaio-dev \
    asciidoc \
    xmlto \
    libgnutls28-dev \
    iproute2 \
    libseccomp-dev \
    curl \
    libbsd-dev \               
    libdrm-dev \                
    libnftables-dev \          
    protobuf-c-compiler          


RUN git clone https://github.com/checkpoint-restore/criu.git /criu && \
    cd /criu && \
    make && \
    make install

#apt-get install inetutils-ping

RUN mkdir /checkpoint-manager
RUN chmod -R 777 /checkpoint-manager

WORKDIR /checkpoint-manager

CMD ["tail", "-f", "/dev/null"]
