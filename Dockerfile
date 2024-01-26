FROM golang:1.18-alpine AS builder
WORKDIR /app/sectask
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o sectask sectask/cmd

FROM alpine
RUN apk update && apk add --no-cache tzdata
WORKDIR /app
COPY --from=builder /app/sectask /app/
ENTRYPOINT ["./sectask"]