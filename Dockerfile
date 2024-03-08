FROM cgr.dev/chainguard/go@sha256:f7c4f59793fce0de9fedd5901ad724c8bb872c428ff23f59cc6eb112dc9a74c2 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
