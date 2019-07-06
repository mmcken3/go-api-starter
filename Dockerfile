FROM golang:1.12.6

WORKDIR /go/src/github.com/mmcken3/go-api-starter
COPY . /go/src/github.com/mmcken3/go-api-starter
RUN go install ./...
CMD ["api"]