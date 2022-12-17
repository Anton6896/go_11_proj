# FROM golang:1.19.3

# ENV TZ=Asia/Jerusalem

# WORKDIR /app
# COPY . .
# RUN go mod download
# RUN go build -o /docker-gs-ping
# RUN rm -rf /app

# CMD [ "/docker-gs-ping" ]

# docker build -t books-go1:ant .
# docker run --rm -it books-go1:ant bash
# docker run --rm books-go1:ant

FROM golang:alpine AS build-env
WORKDIR /src
# RUN apk --no-cache add build-base git mercurial gcc
ADD . .
RUN go mod download
RUN go build -o goapp

FROM alpine
WORKDIR /app
COPY --from=build-env /src/goapp /app/
ENTRYPOINT ./goapp