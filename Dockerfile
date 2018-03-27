FROM golang:1.8

WORKDIR /go/src/testapp
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["testapp"]