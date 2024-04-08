FROM cgr.dev/chainguard/go@sha256:bc4b9e98ca6c4304c93b537c0c8f40715d0b11de2600691700b562fa47f0571c as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
