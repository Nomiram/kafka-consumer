FROM golang:1.19 as builder
# ENV DBADDR="db"
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
# RUN mkdir /app
# WORKDIR /go/bin/app
ENV GOPATH /app/bin
RUN go get kafka-consumer 
RUN CGO_ENABLED=1 GOOS=linux go install -a kafka-consumer
# RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-s -w" -o /app/bin -v ./...

FROM alpine:latest as run
RUN apk --no-cache add ca-certificates
# RUN apk add --no-cache libc6-compat 
RUN apk update && apk add --no-cache libc6-compat &&\
ln -s /lib/libc.musl-x86_64.so.1 /lib/ld-linux-x86-64.so.2
# RUN apk add ld-linux-x86-64
# WORKDIR /root/
RUN addgroup -S app && adduser -S app -G app
COPY --from=builder --chown=app /app /app
RUN chmod +x /app/*
USER app
RUN ls -al /app/bin/

CMD [ "/app/bin/bin/kafka-consumer" ]