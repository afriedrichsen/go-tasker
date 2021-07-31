ARG BASE
FROM golang:1.16.5-alpine3.14 as builder
WORKDIR /src
COPY *.go ./
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download
RUN go build -o /app
# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
# ARG SKAFFOLD_GO_GCFLAGS
# RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /app .

FROM $BASE
COPY --from=builder /app .