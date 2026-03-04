# Build Stage
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o winjur-rest-api .

# Run Stage
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/winjur-rest-api .

EXPOSE 8080
CMD ["./winjur-rest-api"]