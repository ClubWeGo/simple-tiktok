FROM golang:1.19.6 AS builder
WORKDIR $GOPATH/src/videomicro
Add . $GOPATH/src/videomicro
RUN go build .

FROM ubuntu:latest
WORKDIR videomicro
COPY --from=builder /go/src/videomicro/videomicro .
EXPOSE 10001
CMD ["./videomicro"]
#ENTRYPOINT ["./bootstrap.sh"]
