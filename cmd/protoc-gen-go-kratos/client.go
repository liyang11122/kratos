package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
)

const (
	contextPackage       = protogen.GoImportPath("context")
	transportGRPCPackage = protogen.GoImportPath("github.com/go-kratos/kratos/v2/transport/grpc")
	transportHTTPPackage = protogen.GoImportPath("github.com/go-kratos/kratos/v2/transport/http")
)

func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_kratos.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-kratos. DO NOT EDIT")
	g.P("// versions:")
	g.P(fmt.Sprintf("// protoc-gen-go-kratos %s", release))
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	return g
}
