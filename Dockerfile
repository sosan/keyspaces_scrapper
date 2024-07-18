FROM cgr.dev/chainguard/go@sha256:80ebff2d788e242fa5c03a0788f7cf0660d4ab470df95839b3eeb6ced45b3108 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
