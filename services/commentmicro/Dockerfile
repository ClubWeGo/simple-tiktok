FROM golang:1.19.6 AS builder
WORKDIR $GOPATH/src/commentmicro
ADD . .
RUN go get github.com/ClubWeGo/simple-tiktok/services/usermicro@latest
RUN go build .

FROM ubuntu:latest
WORKDIR commentmicro
COPY --from=builder /go/src/commentmicro/commentmicro .
EXPOSE 10004
CMD ["./commentmicro"]
