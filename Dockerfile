FROM golang:1.17

ENV GOPATH=/

COPY . .

RUN go mod download -x
RUN go build -o service ./cmd/main.go

EXPOSE 8080 8080
CMD ["./service"]