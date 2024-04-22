FROM cgr.dev/chainguard/go@sha256:159e1bcd846e4244fdb412845b18e91e34342fe4c518087184b944916f9d03fe as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
