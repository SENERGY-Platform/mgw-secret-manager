<a href="https://github.com/SENERGY-Platform/mgw-secret-manager/actions/workflows/test.yml" rel="nofollow">
        <img src="https://github.com/SENERGY-Platform/mgw-secret-manager/actions/workflows/test.yml/badge.svg" alt="Tests" />
</a>

# Secret Manager
## Run 
Run `make run` to start the dependencies and run the secret manager from the local repository.
Run `make run_with_db` to start a MySQL DB with Docker and run the secret manager from the local repository.
Run `makr run_all_docker` to start the application completely via Docker. The API will be exposed at `8080`

## Tests
Run `make run_test` to start the dependencies and run the tests locally.
Run `make run_test_docker` to to start the dependencies and run the tests inside a Docker container.
