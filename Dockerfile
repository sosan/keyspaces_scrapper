FROM cgr.dev/chainguard/go@sha256:c531ef65c3190e83318cf19b3d6386d3e4be6a66cca2d76bb6ff15bce8ee6a02 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
