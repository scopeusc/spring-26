# builder
FROM golang:1.24 AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# static, stripped binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /server .

# final runtime
FROM gcr.io/distroless/static:nonroot
COPY --from=builder /server /server
USER nonroot
EXPOSE 8080
ENTRYPOINT ["/server"]