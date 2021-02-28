FROM alpine:3.11
RUN mkdir /consul && cd /consul && wget https://releases.hashicorp.com/consul/1.9.3/consul_1.9.3_linux_amd64.zip \
&& apk add unzip && unzip consul_1.9.3_linux_amd64.zip && rm -rf consul_1.9.3_linux_amd64.zip && mv consul /usr/local/bin

EXPOSE 8300
EXPOSE 8301 8301/udp 8302 8302/udp
EXPOSE 8500 8600 8600/udp