FROM ubuntu:20.04
RUN sed -i s@/archive.ubuntu.com/@/mirrors.163.com/@g /etc/apt/sources.list
RUN sed -i s@/security.ubuntu.com/@/mirrors.163.com/@g /etc/apt/sources.list
RUN apt-get clean
RUN apt-get -y update

WORKDIR ~/
Add ./output .
ENTRYPOINT ["./bootstrap.sh"]