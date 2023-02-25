FROM golang:1.19.6 AS builder
WORKDIR $GOPATH/src/simple-tiktok
ADD . .
RUN go get github.com/ClubWeGo/simple-tiktok/services/usermicro@latest
RUN go get github.com/ClubWeGo/simple-tiktok/services/videomicro@latest
RUN go get github.com/ClubWeGo/simple-tiktok/services/favoritemicro@latest
RUN go get github.com/ClubWeGo/simple-tiktok/services/commentmicro@latest
RUN go get github.com/ClubWeGo/simple-tiktok/services/relationmicro@latest
RUN go build .


FROM ubuntu:latest
WORKDIR simple-tiktok
COPY --from=builder /go/src/simple-tiktok/simple-tiktok .
EXPOSE 8888
RUN apt update
RUN apt install -y ffmpeg
CMD ["./simple-tiktok"]
