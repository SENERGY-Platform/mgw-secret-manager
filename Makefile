BINARY_NAME=secret-manager

build:
	go build -o ${BINARY_NAME} main.go

run: 
	docker compose -f deployments/docker-compose.yml up -d --force-recreate db 
	DEV=true ENABLE_ENCRYPTION=false DB_CONNECTION_URL=user:password@tcp\(localhost:3307\)/db go run ./...

run_docker:
	docker compose -f deployments/docker-compose.yml up --build --force-recreate

clean:
	go clean
	docker compose -f test/docker-compose.yml down --remove-orphans
	docker compose -f deployments/docker-compose.yml down --remove-orphans

run_test:
	docker compose -f test/docker-compose.yml up -d --force-recreate db 
	DB_CONNECTION_URL=user:password@tcp\(localhost:3306\)/db go test ./... -coverprofile=cov.xml -coverpkg=./...

run_test_docker:
	docker compose -f test/docker-compose.yml up --build --force-recreate --exit-code-from test