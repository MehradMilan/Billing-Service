# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /billing

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go mod tidy

RUN go build cmd/project.go

EXPOSE 8000

CMD ["./project"]
