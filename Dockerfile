FROM golang:1.23-alpine

WORKDIR /app

COPY cms-api/go.mod cms-api/go.sum ./
RUN go mod tidy

COPY cms-api/ .

RUN go build -o /cms-api

EXPOSE 8080
CMD ["/cms-api"]
