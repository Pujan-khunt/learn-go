FROM golang:1.25

WORKDIR /app

COPY . .

RUN go build -v -o server .

EXPOSE 8080

CMD ["./server"]
