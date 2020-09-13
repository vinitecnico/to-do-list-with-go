FROM golang

RUN mkdir -p /go/src/toDoListGo
WORKDIR /go/src/toDoListGo

ADD . /go/src/toDoListGo

RUN go get github.com/lib/pq

RUN go get -v