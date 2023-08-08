BINARY_NAME=secret-manager

build:
	go build -o ${BINARY_NAME} main.go

run:
	LOGGER_TERMINAL=true DEV=true ENABLE_ENCRYPTION=true DB_CONNECTION_URL=user:password@tcp\(localhost:3307\)/db go run ./...

run_with_db: 
	docker compose -f deployments/docker-compose.yml up -d --force-recreate db 
	SERVER_PORT=80 LOGGER_TERMINAL=true EXPOSE_CONFIDENTIAL_ENDPOINTS=true DEV=true ENABLE_ENCRYPTION=false DB_CONNECTION_URL=user:password@tcp\(localhost:3307\)/db go run ./...

run_all_docker:
	docker compose -f deployments/docker-compose.yml up --build --force-recreate

clean:
	go clean
	docker compose -f test/docker-compose.yml down --remove-orphans
	docker compose -f test/docker-compose.yml rm 
	docker compose -f deployments/docker-compose.yml down --remove-orphans
	docker compose -f deployments/docker-compose.yml rm 

run_test_with_db:
	docker compose -f test/docker-compose.yml up -d --force-recreate db 
	DB_CONNECTION_URL=user:password@tcp\(localhost:3306\)/db go test ./... -coverprofile=cov.xml -coverpkg=./...

run_single_test_with_db:
	docker compose -f test/docker-compose.yml up -d --force-recreate db 
	DB_CONNECTION_URL=user:password@tcp\(localhost:3306\)/db go test -run TestSetKeyWithDisabledEncryption github.com/SENERGY-Platform/mgw-secret-manager/internal/server

run_test_docker:
	docker compose -f test/docker-compose.yml up --build --force-recreate --exit-code-from test

run_load_test:
	locust -f test/locustfile.py