FROM golang:1.16.6-alpine AS builder
RUN apk add --no-cache git
WORKDIR /builer
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /builer/app src/cmd/api/main.go src/cmd/api/wire_gen.go

FROM alpine:3.9
RUN apk add --no-cache ca-certificates
COPY --from=builder /builer/app /app
EXPOSE 8000
CMD ["/app"]