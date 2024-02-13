FROM cgr.dev/chainguard/go@sha256:01bd64717a5b872939b683399a27e2493ea9a5061c84c27a4c99c90fb17179d7 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
