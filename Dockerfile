FROM cgr.dev/chainguard/go@sha256:9bb4a05365c3384c0a0ae925f1a86bd9bddbaed035a899ed7bcf9d7b2dc747ae as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
