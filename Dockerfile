FROM cgr.dev/chainguard/go@sha256:701306bc4743750119b4dc83e54102b493e23f84ee507a98bea477cbb0364522 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
