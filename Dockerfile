FROM jimdo/golang-dep
WORKDIR /go/src/app
COPY . .
RUN make test
RUN make
EXPOSE 8080
CMD ["/go/src/app/aleph-exporter"]
