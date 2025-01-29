mgw-secret-manager
=======

Generate swagger docs:

    swag init -g routes.go -o internal/http_handler/swagger_docs -dir internal/http_handler/standard,internal/http_handler/shared --parseDependency --instanceName standard
    swag init -g routes.go -o internal/http_handler/swagger_docs -dir internal/http_handler/restricted,internal/http_handler/shared --parseDependency --instanceName restricted