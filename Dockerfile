FROM cgr.dev/chainguard/go@sha256:2bacd9b79dbc5ea00fae518411726d6078e612ef8cb8d11f80d200ff67dc45f4 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
