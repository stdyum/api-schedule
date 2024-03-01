FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN rm go.work go.work.sum
RUN go build -o app github.com/stdyum/api-schedule/cmd/schedule

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]