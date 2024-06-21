FROM cgr.dev/chainguard/go@sha256:59b2dced8dc2e8be7b70875ff2ee63b5f32588cb5ebc2c7ed02614f5f31da5ba as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
