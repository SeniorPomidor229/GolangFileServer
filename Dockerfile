FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY /public ./public
RUN go mod download

COPY *go ./

RUN go build -o main main.go

CMD ["./main" ]
