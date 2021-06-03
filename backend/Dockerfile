FROM golang:1.16.4-buster

RUN go get \
    github.com/codegangsta/gin \
    github.com/99designs/gqlgen

CMD [ "make", "dev" ]
