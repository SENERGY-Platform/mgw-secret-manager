BINARY_NAME=secret-manager

build:
	go build -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

test_integration:
	docker compose -f test/docker-compose.yml up -d db 
	TEST_MODE=INTEGRATION DB_CONNECTION_URL=user:password@tcp\(localhost:3306\)/db go test ./... -coverprofile=cov.xml -coverpkg=./...

test_unit:
	TEST_MODE=UNIT go test ./... -coverprofile=cov.xml -coverpkg=./...

dep:
	go mod download

lint:
	golangci-lint run --enable-all