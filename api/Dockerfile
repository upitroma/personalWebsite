FROM golang:1.24-alpine

WORKDIR /app

RUN apk add ffmpeg

COPY . .

RUN go build -o server

EXPOSE 8080

CMD ["./server"]