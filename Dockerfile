FROM golang:1.19.1
RUN  mkdir  -p /app
WORKDIR /app
COPY  . .
RUN    go mod download
RUN    go build app/main.go
ENTRYPOINT ["./main"]







