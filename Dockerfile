FROM cgr.dev/chainguard/go@sha256:68bf4763c4e1b475a9518da80fe597c26b89a6a7221c9464258d57d1543e20e5 as builder

WORKDIR /build

COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app *.go

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
