FROM cgr.dev/chainguard/go@sha256:ec5d9c9807380b5a9b834b7b064164136ceb442e37b85b5c33fe04c3027e2a16 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
