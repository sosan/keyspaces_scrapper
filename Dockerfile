FROM cgr.dev/chainguard/go@sha256:d1dd86fbdfd50c231dc0f8fd9768fbbc31df25baed17e25aab76d24a503a7d44 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
