# nexus-proto

Nexus 公共协议定义、生成代码及业务错误系统。

包含面向客户端和 Agent 的 API 定义（`api/`）、共享数据类型（`shared/`），以及所有业务错误码的唯一定义（`errors/`）。

## 目录结构

```
proto/                  # Protobuf 源文件
├── api/v1/             # 客户端 Connect RPC 服务定义
├── shared/v1/          # 共享数据类型（所有层复用）
├── buf.yaml            # Buf 模块配置
├── buf.gen.go.yaml     # Go message 生成配置
└── buf.gen.connect.yaml# Go Connect handler 生成配置
gen/
└── go/                 # Go 生成代码（已提交，可直接 go get）
    ├── api/v1/
    │   └── apiv1connect/
    └── shared/v1/
errors/                 # 业务错误系统（nxerr）
├── errors.go           # Error 类型、ToConnect()、Is()、IsCode()
├── auth.go             # 认证错误码（1000–1999）
├── conversation.go     # 会话错误码（2000–2999）
├── ...                 # 按领域分文件
```

## 使用

### Go

```go
import (
    apiv1 "github.com/pinealctx/nexus-proto/gen/go/api/v1"
    "github.com/pinealctx/nexus-proto/gen/go/api/v1/apiv1connect"
    sharedv1 "github.com/pinealctx/nexus-proto/gen/go/shared/v1"
    nxerr "github.com/pinealctx/nexus-proto/errors"
)
```

### 错误系统

```go
// 服务端：返回错误，handler 层转换
return nxerr.ErrMessageNotFound
return nil, nxerr.ToConnect(err)

// 客户端：检查错误
if nxerr.Is(err, nxerr.ErrRecallTimeout) { ... }
if nxerr.IsCode(err, nxerr.ErrCardActionExpired.Code) { ... }
```

### TypeScript / Swift / Rust

TS 项目通过 `buf generate` 从 `proto/` 生成代码（见各项目的 `buf.gen.ts.yaml`）。
Swift 和 Rust 通过 `protoc-gen-swift` / `prost` 生成。

## 重新生成代码

需要安装 [buf](https://buf.build/docs/installation)、`protoc-gen-go`、`protoc-gen-connect-go`。

```bash
make generate       # 生成
make clean          # 清理
make lint           # buf lint
```

## 许可证

私有项目。
