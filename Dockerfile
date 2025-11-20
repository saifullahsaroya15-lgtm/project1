FROM golang:1.20-alpine AS build
WORKDIR /app
COPY main.go .
RUN go build -o server main.go
FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/server .
EXPOSE 8009
CMD ["./server"]
