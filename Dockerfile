FROM golang:1.19

WORKDIR /app

RUN go mod init teste

COPY go .

RUN go build -o math

CMD ["./math"]
