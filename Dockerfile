FROM go:1.25.0 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /boards

FROM scratch
COPY --from=builder /boards /boards

EXPOSE 8080

ENTRYPOINT ["/boards"]