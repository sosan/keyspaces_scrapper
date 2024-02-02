FROM cgr.dev/chainguard/go@sha256:5d5abab93c240e0190be34d08c24f32f552cf0aa235ec5a8bbd645071c89c415 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
