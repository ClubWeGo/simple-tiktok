FROM golang:1.19.6 AS builder
WORKDIR $GOPATH/src/relationmicro
ADD . .
RUN go get github.com/ClubWeGo/simple-tiktok/services/usermicro@latest
RUN go build .

FROM ubuntu:latest
WORKDIR relationmicro
COPY --from=builder /go/src/relationmicro/relationmicro .
EXPOSE 10002
CMD ["./relationmicro"]
