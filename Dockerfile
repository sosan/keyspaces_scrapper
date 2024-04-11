FROM cgr.dev/chainguard/go@sha256:b3466202265a95c623a6ded03a8293076555d934dca32e3ef3b0898decfc247d as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
