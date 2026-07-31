FROM --platform=$BUILDPLATFORM golang:1.23 AS builder

ARG TARGETOS=linux
ARG TARGETARCH

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates
CMD ["/app/main"]
