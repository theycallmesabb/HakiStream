FROM --platform=linux/amd64 golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o hakistream .

FROM --platform=linux/amd64 alpine:latest
WORKDIR /app
COPY --from=builder /app/hakistream .
COPY --from=builder /app/static ./static
EXPOSE 8080
CMD ["./hakistream"]