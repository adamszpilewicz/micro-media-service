# Build stage
FROM golang:alpine AS build

ARG PKG_NAME=github.com/johandry/micro-media-service

ADD . /go/src/${PKG_NAME}

RUN cd /go/src/${PKG_BASE}/${PKG_NAME} && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /movie 

# Run stage and microservice image
FROM scratch

COPY --from=build /movie .

EXPOSE 8086

ENTRYPOINT [ "./movie" ]