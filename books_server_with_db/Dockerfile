FROM golang:1.19.3

ENV TZ=Asia/Jerusalem

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /docker-gs-ping

CMD [ "/docker-gs-ping" ]

# docker build -t books-go1:ant .
# docker run --rm -it books-go1:ant bash
# docker run --rm books-go1:ant