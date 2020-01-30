build: vendor aleph-exporter

test: vendor
	go test ./...

vendor:
	dep ensure

aleph-exporter: vendor
	go build ./cmd/...

docker: 
	docker build -t alephexporter .

docker-push: 
	./docker-push.sh

docker-run: docker
	docker run -e ALEPH_HOST=$$ALEPH_HOST -e ALEPH_TOKEN=$$ALEPH_TOKEN -d -p 9090:9090 alephexporter

clean:
	rm -r vendor
	rm aleph-exporter

lint: vendor
	golint
	gosec -exclude=G104 ./...

