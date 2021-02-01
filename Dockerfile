FROM golang
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .


# Build final image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
CMD ["./cmd/web"]