FROM golang:1.19.6 AS builder
WORKDIR $GOPATH/src/favoritemicro
ADD . .
RUN go get github.com/ClubWeGo/simple-tiktok/services/videomicro@latest
RUN go build .

FROM ubuntu:latest AS favoritemicro
WORKDIR favoritemicro
COPY --from=builder /go/src/favoritemicro/favoritemicro .
EXPOSE 10003
CMD ["./favoritemicro"]
