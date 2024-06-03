FROM cgr.dev/chainguard/go@sha256:84bb9f687ca7d6bd70755aabc8687e74370c665457ac666a968228f3f00528d3 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
