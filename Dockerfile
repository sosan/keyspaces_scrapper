FROM cgr.dev/chainguard/go@sha256:526b7fd33ab48b632957197f077e21f895f92c41b7a810eef07fcf288f10dc0d as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
