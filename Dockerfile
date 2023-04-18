FROM golang:alpine
WORKDIR $GOPATH/src/tonguetest
ADD . ./
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
RUN go build -o tonguetest .
EXPOSE 8080
ENTRYPOINT  ["./tonguetest"]