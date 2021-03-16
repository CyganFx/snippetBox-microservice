FROM golang AS builder
COPY . /app
WORKDIR /app
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o myApp ./cmd

FROM alpine:latest
WORKDIR /root
RUN apk --no-cache add ca-certificates
COPY --from=builder /app .

EXPOSE 4011

#CMD in compose