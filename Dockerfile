FROM cgr.dev/chainguard/go@sha256:8a6f881262e5aeeb2a93597dde5aca70cb42e32f474704838191761c96301951 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
