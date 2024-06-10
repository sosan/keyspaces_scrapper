FROM cgr.dev/chainguard/go@sha256:33158972c85a407c195aa3afb8b9c0b3e29d3c84d807bd1575271b8bd02e1d2d as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
