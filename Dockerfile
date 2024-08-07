FROM cgr.dev/chainguard/go@sha256:06a833de4ad9e9b2095de8af3a7097392649ae37528553be73b0855aed425c53 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
