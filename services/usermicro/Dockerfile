FROM golang:1.19.6 AS builder
WORKDIR $GOPATH/src/usermicro
ADD . .
RUN go build .

FROM ubuntu:latest
WORKDIR usermicro
COPY --from=builder /go/src/usermicro/usermicro .
EXPOSE 10000
CMD ["./usermicro"]
