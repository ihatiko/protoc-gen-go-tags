```bash install
go install github.com/ihatiko/protoc-gen-go-tags
```

```bash compile
protoc -I . --go-tags_out protoc ./protoc/*/*.proto
```

```proto
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
