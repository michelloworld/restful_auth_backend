FROM golang:1.17.0-alpine3.14 as builder

LABEL stage=builder

RUN apk update && apk upgrade && \
    apk add --no-cache libc-dev gcc

WORKDIR /go/src/myapp

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o my_app .

# ---

FROM alpine:3.14

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /go/src/myapp/my_app .
COPY --from=builder /go/src/myapp/.env .env


CMD [ "./my_app" ]
