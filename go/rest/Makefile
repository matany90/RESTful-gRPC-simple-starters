init-dist:
	mkdir -p dist
	touch dist/.gitkeep

vendor:
	make clean
	go mod vendor

docker-build:
	docker build . -f ./Dockerfile -t golang-rest-api -q

docker-run:
	docker run -p 50003:50003 golang-rest-api

release:
	make vendor
	goreleaser --snapshot --skip-publish --rm-dist
	make init-dist
	echo "Release complete please check dist directory."

run:
	go run ./cmd/rest/main.go

clean-dist:
	rm -rf ./dist/*
	make init-dist

clean:
	make clean-dist
	rm -rf ./vendor