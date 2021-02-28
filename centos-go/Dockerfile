FROM centos:7
RUN mkdir /go && cd /go && yum install wget -y && wget -q https://golang.google.cn/dl/go1.16.linux-amd64.tar.gz \
&& tar -zxf go1.16.linux-amd64.tar.gz -C /usr/local \
&& rm -rf go1.16.linux-amd64.tar.gz
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
ENV PATH $GOROOT/bin:$PATH
ENV GO111MODULE auto
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/