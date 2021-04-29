FROM golang:1.15.7-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN apk add gcc libc-dev
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o bin/main cmd/main.go

FROM scratch
COPY --from=builder /build/bin/main /app/
WORKDIR /app
CMD ["./main"]
