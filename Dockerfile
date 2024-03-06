FROM cgr.dev/chainguard/go@sha256:a0973622ed0bb3e5a9d5c990598bfb79102b55b155120737a0c30e52d2514be6 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
