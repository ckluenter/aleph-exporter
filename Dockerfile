FROM golang:1.8
WORKDIR /go/src/app
COPY . .
RUN make
EXPOSE 8080
CMD ["/go/src/app/aleph-exporter"]
