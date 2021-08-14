# Go Proto Import Problem

Compilation of `proto` files:

```bash
protoc --proto_path=proto (find proto -name '*.proto') --go_out=plugins=grpc:.
```