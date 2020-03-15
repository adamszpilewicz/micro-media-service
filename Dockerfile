# Build stage
FROM golang:alpine AS build

ARG PKG_NAME=github.com/johandry/micro-media-service

ENV UPX_VER 3.95

ADD . /go/src/${PKG_NAME}
ADD https://github.com/upx/upx/releases/download/v${UPX_VER}/upx-${UPX_VER}-i386_linux.tar.xz /

# Install upx and git to use `go get`
RUN apk --update add git openssh && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/* && \
    tar xf /upx-${UPX_VER}-i386_linux.tar.xz && \
    mv upx-${UPX_VER}-i386_linux/upx /bin/upx

RUN cd /go/src/${PKG_BASE}/${PKG_NAME} && \
    go get github.com/sirupsen/logrus && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o /movie.bin && \
    /bin/upx -k --best --ultra-brute -o /movie /movie.bin

# Run stage and microservice image
FROM scratch AS application

COPY --from=build /movie .

EXPOSE 8086

CMD [ "./movie" ]