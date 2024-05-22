FROM cgr.dev/chainguard/go@sha256:08d316c0f1d271a94cc9fe70e926f531cb3714f87ac9025e6bb64f5e55650541 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
