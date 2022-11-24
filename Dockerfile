FROM golang:1.18-alpine as builder
RUN apk add build-base
RUN apk add --no-cache tzdata
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./cmd/
CMD [ "/app/main" ]