package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/ihatiko/protoc-gen-go-tags/protoc/tags"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	if value := os.Getenv("DEBUG"); value == "true" {
		data, err := os.ReadFile("wire")
		if err != nil {
			panic(err)
		}
		req := &plugin.CodeGeneratorRequest{}
		if err := proto.Unmarshal(data, req); err != nil {
			panic(err)
		}
		gen, err := newGeneratorFromSources(req)
		if err != nil {
			panic(err)
		}
		if gen == nil {
			panic("empty gen")
		}
		for _, f := range gen.Files {
			processing(f)
		}
		return
	}
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			processing(f)
		}
		return nil
	})
}

func newGeneratorFromSources(req *pluginpb.CodeGeneratorRequest, sources ...string) (*protogen.Plugin, error) {
	for _, src := range sources {
		var fd descriptorpb.FileDescriptorProto
		if err := prototext.Unmarshal([]byte(src), &fd); err != nil {
			return nil, err
		}
		req.FileToGenerate = append(req.FileToGenerate, fd.GetName())
		req.ProtoFile = append(req.ProtoFile, &fd)
	}
	return protogen.Options{}.New(req)
}

type Tags map[string]string
type Field map[string]Tags

func processing(file *protogen.File) {
	if !file.Generate {
		return
	}
	fSet := token.NewFileSet()
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filePath := filepath.Join(path, strings.ReplaceAll(*file.Proto.Name, ".proto", ".pb.go"))
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		return
	}
	node, err := parser.ParseFile(fSet, filePath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
		return
	}
	replacer := make(map[string]Field)
	for _, message := range file.Messages {
		msg := message.Desc.Name()
		fields := make(map[string]Tags)
		for _, field := range message.Fields {
			opts, ok := field.Desc.Options().(*descriptorpb.FieldOptions)
			if !ok || opts == nil {
				continue
			}
			tags := collectExtensions(opts)
			if len(tags) > 0 {
				fields[field.GoName] = tags
			}
		}
		replacer[string(msg)] = fields
	}
	if len(replacer) > 0 {
		for k, v := range replacer {
			for _, decl := range node.Decls {
				if genDecs, ok := decl.(*ast.GenDecl); ok {
					for _, specIn := range genDecs.Specs {
						if spec, ok := specIn.(*ast.TypeSpec); ok {
							if spec.Name.Name == k {
								if len(replacer[k]) == 0 {
									continue
								}
								for _, field := range spec.Type.(*ast.StructType).Fields.List {
									if v, ok := v[field.Names[0].Name]; ok {
										tags := strings.Split(field.Tag.Value, " ")
										bf := make(map[string]string)
										for _, i := range tags {
											formatted := strings.TrimSpace(i)
											formatted = strings.Trim(formatted, "`")
											fragments := strings.Split(formatted, ":")
											bf[fragments[0]] = fragments[1]
										}
										for kCurrent, vCurrent := range v {
											bf[kCurrent] = fmt.Sprintf("\"%s\"", vCurrent)
										}
										var result []string
										for k, v := range bf {
											result = append(result, fmt.Sprintf("%s:%s", k, v))
										}
										field.Tag.Value = fmt.Sprintf("`%s`", strings.Join(result, " "))
									}
								}
							}
						}
					}
				}
			}
		}
	}
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fSet, node); err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filePath, buf.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
func collectExtensions(opts *descriptorpb.FieldOptions) map[string]string {
	result := make(map[string]string)
	tags := []*protoimpl.ExtensionInfo{
		annotations.E_Json,
		annotations.E_Param,
		annotations.E_Custom,
		annotations.E_Validate,
		annotations.E_Query,
		annotations.E_Db,
	}
	for _, i := range tags {
		switch extension := proto.GetExtension(opts, i).(type) {
		case string:
			if extension != "" {
				result[string(i.TypeDescriptor().Name())] = extension
			}
		case *annotations.Custom:
			if extension.Value != "" {
				result[extension.Key] = extension.Value
			}
		case []*annotations.Custom:
			for _, cs := range extension {
				if cs.Value != "" {
					result[cs.Key] = cs.Value
				}
			}
		}

	}
	return result
}
