FROM golang:latest

ENV ROOT=/app
WORKDIR ${ROOT}

COPY ./app/main.go ${ROOT}
COPY ./app/go.mod ${ROOT}

RUN go mod tidy