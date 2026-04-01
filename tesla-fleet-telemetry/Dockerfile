FROM golang:1.24-alpine AS builder

WORKDIR /src
COPY launcher.go /src/launcher.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /out/launcher /src/launcher.go

FROM tesla/fleet-telemetry:latest

COPY --from=builder /out/launcher /launcher

CMD ["/launcher"]
