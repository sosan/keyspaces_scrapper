FROM cgr.dev/chainguard/go@sha256:e10e9752d6bd2da2894027a957572e52d6d2bcd8fd29f57c5bdc9978a90211c6 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
