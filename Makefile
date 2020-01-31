build: aleph-exporter

test: 
	go test ./...

aleph-exporter: 
	go build ./cmd/...

docker: 
	docker build -t alephexporter .

docker-push: 
	./docker-push.sh

docker-run: docker
	docker run -e ALEPH_HOST=$$ALEPH_HOST -e ALEPH_TOKEN=$$ALEPH_TOKEN -d -p 9090:9090 alephexporter

clean:
	rm aleph-exporter

lint: 
	golint
	gosec -exclude=G402,G104 ./...

