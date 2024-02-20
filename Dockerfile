FROM cgr.dev/chainguard/go@sha256:e904fb571ca3f60d69e743224f8fc8cfcd0c8faaa29a9307c71e96dde0b70bf1 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
