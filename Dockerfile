FROM cgr.dev/chainguard/go@sha256:74bc9af1d45fd1c8d432a89148c5e413711204636b54ca05197b511bea7a18fb as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
