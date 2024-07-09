FROM cgr.dev/chainguard/go@sha256:0b3fbcaec43485c31b38ecd25d1ad28221f44453d8f66e56103ebc63b3d21c19 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
