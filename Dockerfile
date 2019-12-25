FROM golang:1.13.5-alpine3.10 as builder

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod verify
RUN GOOS=linux GOARCH=arm go build -ldflags="-w -s" -o hello

#FROM alpine:latest
FROM arm64v8/alpine:3.7
COPY --from=builder /app/config /app/config
COPY --from=builder /app/hello /app/hello
EXPOSE 8383/tcp
ENTRYPOINT ["/app/hello"]