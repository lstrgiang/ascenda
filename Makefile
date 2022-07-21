
run:
	go run main.go server --supplier data/suppliers.json

test:
	go test -v ./...

docker/build:
	docker build -t ascenda-server:latest .

docker/run:
	docker run -v data:/data -p 8081:8081 --name ascenda-server ascenda-server
