FROM cgr.dev/chainguard/go@sha256:7713389cd43a75360efff7d62f1c75459bdb47c407d6b3bd3ec50bccc6906646 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
