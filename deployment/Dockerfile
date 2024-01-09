FROM golang:1.21 as builder

WORKDIR /app

COPY . /app

RUN go mod download

RUN CGO_ENABLED=0 go build -o /app/application .

FROM scratch

COPY --from=builder /app/db/migrations /db/migrations
COPY --from=builder /app/application /application
COPY --from=builder /app/.env /.env

EXPOSE 8080

ENTRYPOINT ["/application"]
