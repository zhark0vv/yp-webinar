FROM golang:1.23
WORKDIR /app
COPY app .
RUN go mod tidy
RUN go build -o zap-app ./cmd/logging

CMD ["./zap-app"]
