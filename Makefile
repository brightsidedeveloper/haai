.PHONY: go-bsd react-vite

gobsd:
	protoc --go_out=./go-bsd/internal --proto_path=. api.proto

vite:
	protoc \
		--plugin=protoc-gen-ts_proto=$(shell which protoc-gen-ts_proto) \
		--ts_proto_out=./react-vite/src/api \
		--proto_path=. \
		api.proto
native:
	protoc \
		--plugin=protoc-gen-ts_proto=$(shell which protoc-gen-ts_proto) \
		--ts_proto_out=./react-native/api \
		--proto_path=. \
		api.proto


proto: gobsd vite native

conn:
	psql -h localhost -p 5432 -U admin -d mydb

up:
	docker compose up -d
down:
	docker compose down
clean:
	docker compose down --volumes --remove-orphans