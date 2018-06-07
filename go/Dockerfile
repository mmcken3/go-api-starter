FROM golang:1.9

ARG package

WORKDIR /go/src/github.com/mmcken3/go-api-starter/go
COPY . /go/src/github.com/mmcken3/go-api-starter/go
RUN go-wrapper install "github.com/Clemson-CPSC-4910/s18-fish-findr/go/fisher/cmd/$package"
