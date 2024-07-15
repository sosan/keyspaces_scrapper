FROM cgr.dev/chainguard/go@sha256:2aab312eba021f07c6adf80abe4cb3c34bace048157e3271a2e41991658562ce as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
