FROM cgr.dev/chainguard/go@sha256:54b74a40acfc93d62bd32c72e3afe19bc55e4b2db7baa09d5950f3e5878baf28 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
