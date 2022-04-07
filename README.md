# go-ddd-learning

GoでClean ArchitectureでDDDするサンプル

## Requirements

- Go 1.18 >=
(Workspace modeを使うため)

## Usage

### Gin

```
go run presentation/gin/main.go
```

### gRPC-Gateway

```
go run presentation/grpc-gateway/main.go
```

### API

- `POST /users`
- `PATCH /users/deactivated`
- `POST /tasks`
- `PATCH /tasks/postponed`
