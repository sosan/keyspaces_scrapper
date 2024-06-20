FROM cgr.dev/chainguard/go@sha256:ae3a8b7efe98b1fa1faa637e11bba8ff4de33981e76b49eb8cad75520b8ae85d as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
