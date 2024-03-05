```bash install
go install github.com/ihatiko/protoc-gen-go-tags
```

```bash compile
protoc -I . --go-tags_out protoc ./protoc/*/*.proto
```
