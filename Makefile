ENVOY=$(shell find ../common-proto/envoy -name "*.proto")
UDPA=$(shell find ../common-proto/udpa -name "*.proto")
APIS=${ENVOY} ${UDPA}

all:	rpc grpc connect

rpc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	protoc ${APIS} \
		--proto_path='../common-proto' \
		--go_opt='module=github.com/agentio/common-go' \
		--go_opt=Mudpa/annotations/status.proto=github.com/agentio/common-go/udpa \
		--go_opt=Mudpa/annotations/versioning.proto=github.com/agentio/common-go/udpa \
		--go_opt=Mudpa/annotations/migrate.proto=github.com/agentio/common-go/udpa \
		--go_opt=Mudpa/annotations/security.proto=github.com/agentio/common-go/udpa \
		--go_opt=Mudpa/annotations/sensitive.proto=github.com/agentio/common-go/udpa \
		--go_opt=Mudpa/data/orca/v1/orca_load_report.proto=github.com/agentio/common-go/udpa \
		--go_opt=Mudpa/service/orca/v1/orca.proto=github.com/agentio/common-go/udpa \
		--go_opt=Mudpa/type/v1/typed_struct.proto=github.com/agentio/common-go/udpa \
		--go_out='.'

grpc:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	protoc ${APIS} \
		--proto_path='../common-proto' \
		--go-grpc_opt='module=github.com/agentio/common-go' \
		--go-grpc_opt=Mudpa/annotations/status.proto=github.com/agentio/common-go/udpa \
		--go-grpc_opt=Mudpa/annotations/versioning.proto=github.com/agentio/common-go/udpa \
		--go-grpc_opt=Mudpa/annotations/migrate.proto=github.com/agentio/common-go/udpa \
		--go-grpc_opt=Mudpa/annotations/security.proto=github.com/agentio/common-go/udpa \
		--go-grpc_opt=Mudpa/annotations/sensitive.proto=github.com/agentio/common-go/udpa \
		--go-grpc_opt=Mudpa/data/orca/v1/orca_load_report.proto=github.com/agentio/common-go/udpa \
		--go-grpc_opt=Mudpa/service/orca/v1/orca.proto=github.com/agentio/common-go/udpa \
		--go-grpc_opt=Mudpa/type/v1/typed_struct.proto=github.com/agentio/common-go/udpa \
		--go-grpc_out='.'

connect:
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
	protoc ${APIS} \
		--proto_path='../common-proto' \
		--connect-go_opt='module=github.com/agentio/common-go' \
		--connect-go_opt=Mudpa/annotations/status.proto=github.com/agentio/common-go/udpa \
		--connect-go_opt=Mudpa/annotations/versioning.proto=github.com/agentio/common-go/udpa \
		--connect-go_opt=Mudpa/annotations/migrate.proto=github.com/agentio/common-go/udpa \
		--connect-go_opt=Mudpa/annotations/security.proto=github.com/agentio/common-go/udpa \
		--connect-go_opt=Mudpa/annotations/sensitive.proto=github.com/agentio/common-go/udpa \
		--connect-go_opt=Mudpa/data/orca/v1/orca_load_report.proto=github.com/agentio/common-go/udpa \
		--connect-go_opt=Mudpa/service/orca/v1/orca.proto=github.com/agentio/common-go/udpa \
		--connect-go_opt=Mudpa/type/v1/typed_struct.proto=github.com/agentio/common-go/udpa \
                --connect-go_out='.'

