FROM golang:1.17.7-alpine3.15 AS builder
WORKDIR /app
COPY . .
ARG CGO_ENABLED=0
ARG GOOS=linux
RUN go build -o build/main ./main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder ./app/build .
COPY --from=builder ./app/data ./data
CMD ["./main"]