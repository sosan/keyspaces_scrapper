FROM cgr.dev/chainguard/go@sha256:7db365e06f91f903cfc7f22f62316525e5a4666708ed0fd9a70fc96c9b442af8 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
