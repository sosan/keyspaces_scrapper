FROM cgr.dev/chainguard/go@sha256:4d51574ef33b4edc57a22da062fe335a500eda30a1f1315cb39b4977bf2aef5f as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
