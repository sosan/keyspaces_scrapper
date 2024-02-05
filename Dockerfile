FROM cgr.dev/chainguard/go@sha256:f29f9a9808d26acbbbfd71f0312c21c8bd89a1184971a5eee188f1070e148bef as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
