FROM centos:centos7

ENV container docker

RUN yum -y update; yum clean all
RUN yum -y install git; yum clean all

# set golang
ENV VERSION 1.9
ENV FILE go$VERSION.linux-amd64.tar.gz
ENV URL https://storage.googleapis.com/golang/$FILE
ENV MD5 84b51207cdd227248ec4e60cc5e2fa79

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN set -eux &&\
  yum -y clean all &&\
  curl -OL $URL &&\
  echo "$MD5  $FILE" | md5sum -c - &&\
  tar -C /usr/local -xzf $FILE &&\
  rm $FILE &&\
  mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN go get -u github.com/constabulary/gb/...

# set project
ENV projectname easysender

COPY . /go/src/$projectname
WORKDIR /go/src/$projectname

RUN gb build

CMD ["./bin/easysender", "-c", "./conf/easysender.yaml"]
