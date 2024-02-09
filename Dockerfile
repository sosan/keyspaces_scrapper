FROM cgr.dev/chainguard/go@sha256:73edce09ed89d2c30548cc6eb163fbfa2dbbce16754a55e69a6f68c9020a3672 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
