FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./review-service/main.go

EXPOSE 8081

CMD ["./main"]