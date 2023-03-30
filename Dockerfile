FROM golang:1.20
RUN mkdir /app
ADD . /app
RUN go env -w GOPROXY="https://goproxy.cn,direct"
WORKDIR /app
RUN make
CMD ["./main"]


