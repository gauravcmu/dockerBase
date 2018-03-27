FROM golang:1.8

RUN go get github.com/gin-gonic/gin && \
go get github.com/gin-contrib/sse

WORKDIR /go/src/github.com/dockerBase
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["dockerBase"]