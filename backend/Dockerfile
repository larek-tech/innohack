# Builder

FROM golang:1.23-alpine AS builder
RUN apk add --update make git curl

ARG MODULE_NAME=github.com/larek-tech/innohack/backend

RUN apk upgrade --update-cache --available && \
    apk add openssl && \
    rm -rf /var/cache/apk/*

WORKDIR /home/${MODULE_NAME}

COPY . /home/${MODULE_NAME}

RUN go build \
		-o ./bin/main$(shell go env GOEXE) ./cmd/server/main.go


# Service

FROM alpine:latest as server
ARG MODULE_NAME=github.com/larek-tech/innohack/backend
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /home/${MODULE_NAME}/bin/main .
COPY --from=builder /home/${MODULE_NAME}/static static/

RUN chown root:root main