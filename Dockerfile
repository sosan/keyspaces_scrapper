FROM cgr.dev/chainguard/go@sha256:1d4242596ec0bd9759f3141b188a42081c0064173e8ed2364b5dbc4db87419ac as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
