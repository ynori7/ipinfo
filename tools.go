//+build tools

package main

//go:generate env GO111MODULE=on 
//TODO install protoc
//go:generate go install github.com/golang/protobuf/protoc-gen-go
//go:generate go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

//go:generate protoc -I=api --go_out=api --doc_out=./docs --doc_opt=html,ip.html api/ipinfo.proto
//go:generate protoc -I=api --go_out=api --doc_out=./docs --doc_opt=html,errors.html api/errors.proto
