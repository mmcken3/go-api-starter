FROM golang:1.11.5

WORKDIR /go/src/github.com/mmcken3/go-api-starter
COPY . /go/src/github.com/mmcken3/go-api-starter
RUN go install ./...
CMD ["api"]