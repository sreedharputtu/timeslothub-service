FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /timeslothub-service main.go

EXPOSE 8080

CMD ["/timeslothub-service"]
