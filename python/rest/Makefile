run:
	python3 main.py

docker-build:
	docker build . -f ./Dockerfile -t python-rest-api -q

docker-run:
	docker run -p 50005:50005 python-rest-api