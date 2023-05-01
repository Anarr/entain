# Build container
FROM golang:1.19 as builder

WORKDIR /
COPY . .

ENV GO111MODULE=on
RUN apt-get update

RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o entain .

FROM alpine
# Run container
# Application configuration
RUN apk update && apk add bash
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /entain /app/

# Application configuration
COPY --from=builder /internal/config/config.json /app/
COPY --from=builder /internal/config/config.json /app/internal/config/
COPY ./internal/database/migrations /app/internal/database/migrations

WORKDIR /app
EXPOSE 8888

CMD ["./entain"]
