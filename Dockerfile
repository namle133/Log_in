FROM golang:1.18.2-alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
EXPOSE 8000

RUN go build -o main ./cmd/main.go

CMD ["/app/main"]