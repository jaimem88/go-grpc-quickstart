# Path to proto files
PROTO_PATH=./proto
PROTO_FILE=$(PROTO_PATH)/echo
DOCKER_IMAGE=echo-protoc

.PHONY: proto
proto: build
	# Generates Go code
	docker run -v $(PWD)/:/tmp -w /tmp $(DOCKER_IMAGE) protoc \
		-I. \
		--lint_out=. \
		--go_out=plugins=grpc,paths=source_relative:. \
		$(PROTO_FILE).proto

.PHONY: clean
clean:
	# Clean up previously generated Go files
	rm $(PROTO_FILE).pb.* || true

.PHONY: build
build: clean
	docker build -t $(DOCKER_IMAGE) .


.PHONY: client
client:
	go run ./cmd/client/main.go

.PHONY: server
server:
	go run ./cmd/server/main.go
	