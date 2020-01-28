build: vendor aleph-exporter

test: vendor
	go test ./...

vendor:
	dep ensure

aleph-exporter: vendor
	go build ./cmd/...

docker: 
	docker build -t aleph-exporter .

docker-push: 
	./docker-push.sh

docker-run: docker
	docker run -d -p 9090:9090 aleph-exporter

clean:
	rm -r vendor
	rm aleph-exporter

lint: vendor
	golint
	gosec -exclude=G104 ./...

