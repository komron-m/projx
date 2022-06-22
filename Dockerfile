# BUILD STAGE
FROM golang:1.18.3-alpine3.16 as builder
WORKDIR /opt
COPY . .
RUN go build -o main ./cmd/awesome_app

# RUN STAGE
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /opt/main .
COPY .env .
COPY wait-for.sh .

EXPOSE 4000

ENTRYPOINT [ "/app/wait-for.sh", "postgres:5432", "--" ]
CMD ["/app/main"]
