# build stage
FROM golang:1.24.3-alpine AS builder
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
# build the binary from cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd

# final stage
FROM gcr.io/distroless/base-debian11
COPY --from=builder /src/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]
