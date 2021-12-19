FROM golang:1.17-alpine AS go_builder

WORKDIR /app

COPY . .

RUN apk add gcc musl-dev && \
    go mod tidy && \
    go get -d -u && \
    go test -v ./... -cover && \
    go build -o paackgo


# Building to finish docker container
FROM alpine:3.14

WORKDIR /app

COPY --from=go_builder /app/paackgo .
COPY --from=go_builder /app/.env .

ENTRYPOINT [ "./paackgo" ]
