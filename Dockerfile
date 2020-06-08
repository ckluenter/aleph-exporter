FROM golang:1.13 as builder
WORKDIR /app/
COPY . .
ENV CGO_ENABLED=0
RUN make
EXPOSE 8080
CMD ["/app/aleph-exporter"]

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /app/aleph-exporter /app/aleph-exporter
CMD ["/app/aleph-exporter"]
