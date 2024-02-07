FROM cgr.dev/chainguard/go@sha256:47411999e142a53832717f28c7a6ece4d522152c12026ce0311fc4192088ea14 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
