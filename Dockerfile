FROM golang:1.12
WORKDIR /app/
COPY . .
RUN make test
RUN make
EXPOSE 8080
CMD ["/app/aleph-exporter"]
