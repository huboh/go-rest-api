FROM  golang:1.22.3-alpine AS builder

WORKDIR /app

COPY . .

RUN apk update && apk add --no-cache make

RUN go mod download

RUN make build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/main /app/bin/main

# empty env file so godotenv pkg don't error
RUN touch .env

CMD ["/app/bin/main"]