FROM cgr.dev/chainguard/go@sha256:959c0502d67d23823ff71caa967a17563f30e03fcb93dee5c85a942e3468beeb as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
