BINARY_NAME=secret-manager

build:
 go build -o ${BINARY_NAME} main.go

run: build
 ./${BINARY_NAME}

clean:
 go clean
 rm ${BINARY_NAME}

test:
 TEST_MODE=INTEGRATION DB_CONNECTION_URL=postgres://user:password@localhost:5432/db go test ./...

dep:
 go mod download

vet:
 go vet

lint:
 golangci-lint run --enable-all