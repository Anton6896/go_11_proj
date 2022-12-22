FROM golang:alpine AS build-env

WORKDIR /app
ADD . .
RUN go mod download
RUN go build -o goapp

FROM alpine

ENV TZ=Asia/Jerusalem

WORKDIR /app
COPY --from=build-env /app/goapp /app/
ENV ZONEINFO=/zoneinfo.zip
ENTRYPOINT ./goapp


# docker build -t books-go1:ant .
# docker run --rm -it books-go1:ant bash
# docker run --rm books-go1:ant