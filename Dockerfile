from docker.io/golang:1.26.5-alpine3.24 as builder

WORKDIR /app

COPY . .

RUN go mod tidy \
 && go build -o /app/evaluation .

from docker.io/golang:1.26.5-alpine3.24

WORKDIR /app

RUN addgroup -S evaluation-service && adduser -S evaluation-service -G evaluation-service
USER evaluation-service

COPY --from=builder --chown=evaluation-service /app/evaluation ./evaluation

CMD ["./evaluation"]
