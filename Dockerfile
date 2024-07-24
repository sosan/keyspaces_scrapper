FROM cgr.dev/chainguard/go@sha256:28343b6ffe88cbfda301d70984a49d4f24bbad56c055e20376524e116bd6e2b4 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
