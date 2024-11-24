FROM golang:alpine AS builder
ENV CGO_ENABLED 0
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -ldflags="-s -w" -o parser1c .
FROM scratch
WORKDIR /build
COPY --from=builder /build/parser1c /build/parser1c
CMD ["./parser1c"]