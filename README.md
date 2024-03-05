```bash install
go install github.com/ihatiko/protoc-gen-go-tags
```

```bash compile
protoc -I . --go-tags_out protoc ./protoc/*/*.proto
```

```proto syntax
message Example {
  string Field1 = 1 [
    (tags.json) = "abs",
    (tags.validate) = "uuid",
    (tags.custom) = {
      key: "schema",
      value: "some-schema"
    },
    (tags.custom) = {
      key: "schema",
      value: "some-schema"
    },
    (tags.custom) = {
      key: "schema",
      value: "some-schema"
    }
  ];
}
```

```go result
type Example struct {
	state		protoimpl.MessageState
	sizeCache	protoimpl.SizeCache
	unknownFields	protoimpl.UnknownFields

	Field1	string	`protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"abs" schema:"some-schema" validate:"uuid"`
}
```
