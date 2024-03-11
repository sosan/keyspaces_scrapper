FROM cgr.dev/chainguard/go@sha256:f997310500262e3c152f94683da9b796f9ec6501ceace4cd4842bfc52e49a12c as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
