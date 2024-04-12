FROM cgr.dev/chainguard/go@sha256:a06a462f22445088e8bbb4478dedf83228af0db9003cd4f4cde5981694bc3d3d as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
