LOCAL_BIN:=$(CURDIR)/bin

install-go-deps:
	go get github.com/pressly/goose/v3/cmd/goose@latest
	go get github.com/jackc/pgx/v5/stdlib
	go get github.com/Masterminds/squirrel
	go get github.com/stretchr/testify/require
	go get github.com/ilyakaznacheev/cleanenv
	go get -u github.com/golang-jwt/jwt/v5

generate-auth-proto:
	protoc --go_out=pkg/auth_v1 --go_opt=paths=import \
	--go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=import \
	api/auth_v1/auth_grpc.proto

generate-user-proto:
	protoc --go_out=pkg/user_v1 --go_opt=paths=import \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=import \
	api/user_v1/user_grpc.proto