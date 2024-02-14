FROM cgr.dev/chainguard/go@sha256:d26c152c7922fdf0123327795f84d727a81035171e2a2aba614ee6f5faa254b4 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
