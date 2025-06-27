FROM golang:1.24.3-bookworm AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 go build -v -o /app/app .

FROM gcr.io/distroless/static-debian12:latest

COPY --from=builder /app/app /
CMD ["/app"]