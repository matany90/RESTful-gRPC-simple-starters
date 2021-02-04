run:
	python3 main.py

docker-build:
	docker build . -f ./Dockerfile -t python-grpc-api -q

docker-run:
	docker run -p 50006:50006 python-grpc-api

compile-proto:
	bash compileproto.sh