FROM cgr.dev/chainguard/go@sha256:4f11a0dfbd73832405bc3f611e53b4dbd61a1d1d23d205f2665cabfbd295a109 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
