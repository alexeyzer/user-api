FROM golang:1.17

ENV GOPATH=/

COPY . .

RUN go mod download -x
RUN go build -o service ./cmd/main.go

EXPOSE 8079 8080 8082 6379 5432
CMD ["./service"]