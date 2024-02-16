FROM cgr.dev/chainguard/go@sha256:e56a25c9b9cd658714601bf00f1c243dd16494cf6b974428cc3902503fa48a86 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
