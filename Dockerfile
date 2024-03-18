FROM cgr.dev/chainguard/go@sha256:203b27e4c8681c26648a40de5db633b1af70d4e7a7512575038ec103ee4a91a2 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
