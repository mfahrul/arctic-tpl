# Builder
FROM golang:1.12.8-alpine3.10 as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY ./App .

RUN GOOS=linux go build -ldflags="-s -w" -tags=main -o auth_service

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/* && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 8080

COPY --from=builder /app/auth_service /app

CMD /app/auth_service