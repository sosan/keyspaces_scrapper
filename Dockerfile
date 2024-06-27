FROM cgr.dev/chainguard/go@sha256:b40841caf313a4baa43873c6cf3efc50d3e8baa1a02a15b0ac8786c07b28e033 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
