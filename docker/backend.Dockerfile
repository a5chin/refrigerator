FROM golang:1.20

RUN mkdir /go/src/ref
WORKDIR /go/src/ref

RUN go install github.com/cosmtrek/air@v1.43.0 \
    && go install github.com/swaggo/swag/cmd/swag@latest
