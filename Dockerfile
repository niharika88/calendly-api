
FROM golang:1.23-alpine AS build

RUN apk add --no-cache git make \
    && apk add aws-cli --no-cache

WORKDIR /app

COPY . .

RUN make build

EXPOSE 2090

CMD ["./build/calendly-api"]