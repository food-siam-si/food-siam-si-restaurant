dev:
	air server -c air.toml

protoc:
	protoc --go_out=. \
        --go_opt=paths=source_relative \
        --go-grpc_out=. \
        --go-grpc_opt=paths=source_relative \
        internal/handlers/proto/*.proto