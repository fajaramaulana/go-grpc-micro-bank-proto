GO_MODULE := github.com/fajaramaulana/go-grpc-micro-bank-proto

.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
	if exist "protogen" rd /s /q protogen
	mkdir protogen\go
else
	rm -fR ./protogen 
	mkdir -p ./protogen/go
endif


.PHONY: protoc-go
protoc-go:
	rm -f protogen/go/bank/*.go
	protoc --proto_path=proto \
 	--go_out=. --go_opt=module=${GO_MODULE} \
  	--go-grpc_out=. --go-grpc_opt=module=${GO_MODULE} \
  	bank/service.proto \
  	bank/type/account.proto \
  	bank/type/exchange.proto \
  	bank/type/transfer.proto \
  	bank/type/transaction.proto \
	resilliency/resilliency.proto \

.PHONY: build
build: clean protoc-go


.PHONY: pipeline-init
pipeline-init:
	sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


.PHONY: pipeline-build
pipeline-build: pipeline-init build