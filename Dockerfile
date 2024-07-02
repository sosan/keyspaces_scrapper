FROM cgr.dev/chainguard/go@sha256:7d7681e677fdf426b98a6772e4bdabff1f0f4bda6a816602e86498b0baefed58 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
