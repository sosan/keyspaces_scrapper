FROM cgr.dev/chainguard/go@sha256:9aa4a854b43f17f60257be559dd2faed470f38a6b0d78d76f3fda47a08bc024a as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
