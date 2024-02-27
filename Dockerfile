FROM cgr.dev/chainguard/go@sha256:1723273a83351c469105e23227a2be725665e88e00e1f53ed2ffa02de46684dd as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
