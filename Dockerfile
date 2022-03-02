FROM golang:1.17-alpine
ENV TIMEZONE Asia/Bangkok

WORKDIR /src
COPY . .

RUN apk add build-base
RUN go mod download
RUN go mod vendor