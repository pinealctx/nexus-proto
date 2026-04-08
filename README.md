# nexus-proto

Nexus 公共协议定义及生成代码。

包含面向客户端和 Agent 的 API 定义（`api/`）以及共享数据类型（`shared/`）。

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
```

## 使用

### Go

```go
import (
    apiv1 "github.com/pinealctx/nexus-proto/gen/go/api/v1"
    "github.com/pinealctx/nexus-proto/gen/go/api/v1/apiv1connect"
    sharedv1 "github.com/pinealctx/nexus-proto/gen/go/shared/v1"
)
```

### 重新生成代码

需要安装 [buf](https://buf.build/docs/installation)、`protoc-gen-go`、`protoc-gen-connect-go`。

```bash
# 生成
make generate

# 清理
make clean
```

## Lint

```bash
make lint
```
