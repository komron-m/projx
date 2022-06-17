# BUILD STAGE
FROM golang:1.18.3-alpine3.16 as builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/awesome_app

# RUN STAGE
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 4000
CMD ["/app/main"]
