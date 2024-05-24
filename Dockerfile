FROM cgr.dev/chainguard/go@sha256:38db527239c4b33c6ebd1751007041e33cdf81de45f2195b1a60816eb351def1 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
