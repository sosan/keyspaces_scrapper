FROM cgr.dev/chainguard/go@sha256:64075ba4901952c84c67980bbc280f6aeb3a0e58f6e5e32021b0d1531e5fea63 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
