FROM cgr.dev/chainguard/go@sha256:5e6348886cf06f56f6e9a78cd6fb0cd3e4b4d896b27653f448ca722b528c460a as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
