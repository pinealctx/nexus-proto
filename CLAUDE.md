# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Canonical source for all public Nexus protobuf definitions (api + shared) and generated Go code. All Go projects import this module; all TS projects generate code from `proto/`.

Also contains the `errors` package — the single source of truth for all Nexus business error codes, names, and Connect RPC status code mappings.

## Commands

```bash
make generate       # Generate Go code from proto definitions
make clean          # Remove all generated code
make lint           # buf lint
```

Requires: `buf` CLI, `protoc-gen-go`, `protoc-gen-connect-go`.

## Structure

```
proto/
  api/v1/           Client-facing Connect RPC service definitions
  shared/v1/        Shared data types, enums, error codes, options
  buf.yaml          Buf module config (deps: bufbuild/protovalidate)
  buf.gen.go.yaml   Go message generation config
  buf.gen.connect.yaml  Go Connect handler generation config
gen/
  go/
    api/v1/         Generated Go types + apiv1connect/ handlers
    shared/v1/      Generated Go shared types
errors/             Business error system (nxerr)
  errors.go         Error type, ToConnect(), Is(), IsCode()
  auth.go           Auth error codes (1000–1999)
  contact.go        Contact error codes
  conversation.go   Conversation error codes (2000–2999)
  group.go          Group error codes
  ...               Domain-segmented error definitions
```

## Go Import Paths

```go
import (
    apiv1 "github.com/pinealctx/nexus-proto/gen/go/api/v1"
    "github.com/pinealctx/nexus-proto/gen/go/api/v1/apiv1connect"
    sharedv1 "github.com/pinealctx/nexus-proto/gen/go/shared/v1"
    nxerr "github.com/pinealctx/nexus-proto/errors"
)
```

## Error System (nxerr)

All Nexus business errors are defined here. Server and client code both import this package.

```go
// Server: return error, handler converts
return nxerr.ErrMessageNotFound
return nil, nxerr.ToConnect(err)

// Client: check error
if nxerr.Is(err, nxerr.ErrRecallTimeout) { ... }
if nxerr.IsCode(err, nxerr.ErrCardActionExpired.Code) { ... }
```

## Proto Conventions

- IDs: `int32` for entity IDs, `int64` for message_id/conversation_id. Timestamps: `int64` Unix ms.
- Enums: `UPPER_SNAKE_CASE` with `{ENUM_NAME}_UNSPECIFIED = 0`.
- No `google.protobuf.Empty` — use named empty messages. No `google.protobuf.Timestamp`.
- Request validation via `buf.validate` annotations.
- Sensitive fields: annotate with `(shared.v1.sensitive) = true`.
- All comments in English. Every message, field, enum value needs a leading comment.
- Error codes in `shared/v1/error_codes.proto` with domain-segmented ranges.

## Buf Lint Rules

STANDARD + COMMENTS + UNARY_RPC. Exceptions: PACKAGE_VERSION_SUFFIX, RPC_NO_SERVER_STREAMING.

## Language Policy

Conversational replies and spec documents in Simplified Chinese. All code and proto comments in English.
