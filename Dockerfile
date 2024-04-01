FROM cgr.dev/chainguard/go@sha256:c26ed7fe56d6752a79c8165bc057472f95054d65e55ce450f247d85cc58fafaa as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
