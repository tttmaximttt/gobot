FROM golang:1.13.5-alpine3.10 as builder

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod verify
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o hello
FROM golang:1.13.5-alpine3.10
COPY --from=builder /app/config /app/config
# Copy our static executable
COPY --from=builder /app/hello /app/hello
RUN ls /app/config
EXPOSE 8383/tcp
ENTRYPOINT ["/app/hello"]