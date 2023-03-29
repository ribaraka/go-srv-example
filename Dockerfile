ARG GO_VERSION=1.18.2
FROM golang:${GO_VERSION} AS builder

ARG ARCH=amd64
ARG GO111MODULE=on

WORKDIR /accounts

COPY go.mod go.sum ./

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build \
		-o ./app \
		-mod=vendor \
		-a ./cmd

FROM alpine:latest
COPY --from=builder /accounts/app /app

ENTRYPOINT ["/app"]
