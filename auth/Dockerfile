FROM golang:1.16.4-buster

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*

RUN go get github.com/codegangsta/gin
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

CMD [ "make", "dev" ]
