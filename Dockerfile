FROM golang AS builder
COPY . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o myApp ./cmd/web

FROM alpine:latest
WORKDIR /root
RUN apk --no-cache add ca-certificates
COPY --from=builder /app /root/

#ubral CMD["./myApp"] v compose (tipo po best practice)tcr