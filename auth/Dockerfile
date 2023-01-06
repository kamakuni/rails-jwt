FROM golang:latest

RUN apt-get update -qq && apt-get install -y git
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]