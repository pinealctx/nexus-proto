package errors

import "connectrpc.com/connect"

// Push error codes (9000-9999). Reference: docs/designs/d12-error-codes.md.
var (
	ErrInvalidPushToken        = New(connect.CodeInvalidArgument, 9001, "INVALID_PUSH_TOKEN")
	ErrUnsupportedPushPlatform = New(connect.CodeInvalidArgument, 9002, "UNSUPPORTED_PUSH_PLATFORM")
	ErrPushTokenNotFound       = New(connect.CodeNotFound, 9003, "PUSH_TOKEN_NOT_FOUND")
)
