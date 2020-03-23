FROM golang:1.14-alpine3.11 as builder
WORKDIR server
COPY server.go go.mod go.sum ./
RUN go build -o /server .

FROM alpine:3.11
CMD ["./server"]
COPY --from=builder /server .