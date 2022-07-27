# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build cmd/project.go

EXPOSE 8000

CMD [ "go", "run", "project.go" ]
