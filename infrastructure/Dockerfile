FROM golang:latest

WORKDIR /app

COPY . /app

RUN go mod download

ENV MONGO_HOST=localhost
ENV MONGO_PORT=27017

EXPOSE 8080

CMD ["go", "run", "cmd/main.go"]
