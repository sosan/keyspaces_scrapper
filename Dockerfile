FROM cgr.dev/chainguard/go@sha256:b135b47a56d2a4a6ad89f076181ae6ae80030cfa701024312cf9a5b0578568a3 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
