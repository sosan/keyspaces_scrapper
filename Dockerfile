FROM cgr.dev/chainguard/go@sha256:c4e6bc328e72849591edac7ad59e239ab329848ab9ad3d47bafe627f27720103 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
