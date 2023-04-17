FROM golang:latest

WORKDIR /app
COPY . .

RUN go build src/agent/server.go

EXPOSE 8000
CMD ["./server"]