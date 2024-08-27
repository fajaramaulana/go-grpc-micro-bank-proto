# README

Welcome to the `go-grpc-micro-bank-proto` repository!

This repository contains the protocol buffer definitions for the Go gRPC micro-bank project.

## Prerequisites

Before getting started, make sure you have the following installed:

- Go (version 1.16 or higher)
- Protocol Buffers (version 3.0 or higher)

## Installation

To install the necessary dependencies, run the following command:

```shell
go get google.golang.org/protobuf/cmd/protoc-gen-go
```

## Usage

To generate the Go code from the protocol buffer definitions, run the following command:

```shell
protoc --go_out=. *.proto
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

