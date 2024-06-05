FROM cgr.dev/chainguard/go@sha256:ffce5cfd74e0475306301a2c62c6bde0b2b337f5a7ebfa62b7e89b3c85ddc5db as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
