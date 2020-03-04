FROM golang:1.14.0 AS builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go # CGO_ENABLED=0 needed to statically link binary

FROM alpine:3.9.5
COPY --from=builder /app/main .
CMD ["./main"]

