package errors

import "connectrpc.com/connect"

// Auth error codes (1000-1999). Reference: docs/designs/d12-error-codes.md.

// Authentication (1001-1099).
var (
	ErrInvalidToken        = New(connect.CodeUnauthenticated, 1001, "INVALID_TOKEN")
	ErrTokenExpired        = New(connect.CodeUnauthenticated, 1002, "TOKEN_EXPIRED")
	ErrRefreshInvalid      = New(connect.CodeUnauthenticated, 1003, "REFRESH_TOKEN_INVALID")
	ErrRefreshReused       = New(connect.CodeUnauthenticated, 1004, "REFRESH_TOKEN_REUSED")
	ErrDeviceLimitExceeded = New(connect.CodeResourceExhausted, 1005, "DEVICE_LIMIT_EXCEEDED")
	ErrRefreshRateLimited  = New(connect.CodeResourceExhausted, 1006, "REFRESH_RATE_LIMITED")
)

// Gateway connection (1007-1099).
var (
	ErrInvalidFrame       = New(connect.CodeInvalidArgument, 1007, "INVALID_FRAME")
	ErrConnectionReplaced = New(connect.CodeAborted, 1008, "CONNECTION_REPLACED")
)

// Verify code (1101-1199).
var (
	ErrVerifyCodeRateLimited     = New(connect.CodeResourceExhausted, 1101, "VERIFY_CODE_RATE_LIMITED")
	ErrVerifyCodeDailyLimit      = New(connect.CodeResourceExhausted, 1102, "VERIFY_CODE_DAILY_LIMIT")
	ErrVerifyCodeSendFailed      = New(connect.CodeInternal, 1103, "VERIFY_CODE_SEND_FAILED")
	ErrVerifyTokenExpired        = New(connect.CodeNotFound, 1104, "VERIFY_TOKEN_EXPIRED")
	ErrVerifyCodeInvalid         = New(connect.CodeInvalidArgument, 1105, "VERIFY_CODE_INVALID")
	ErrVerifyCodeTooManyAttempts = New(connect.CodeResourceExhausted, 1106, "VERIFY_CODE_TOO_MANY_ATTEMPTS")
)

// Password (1201-1299).
var (
	ErrInvalidPassword    = New(connect.CodeUnauthenticated, 1201, "INVALID_PASSWORD")
	ErrAccountLocked      = New(connect.CodeResourceExhausted, 1202, "ACCOUNT_LOCKED")
	ErrPasswordTooWeak    = New(connect.CodeInvalidArgument, 1203, "PASSWORD_TOO_WEAK")
	ErrResetTokenInvalid  = New(connect.CodeNotFound, 1204, "RESET_TOKEN_INVALID")
	ErrPasswordNotSet     = New(connect.CodeFailedPrecondition, 1205, "PASSWORD_NOT_SET")
	ErrSamePassword       = New(connect.CodeInvalidArgument, 1206, "SAME_PASSWORD")
	ErrPasswordAlreadySet = New(connect.CodeFailedPrecondition, 1207, "PASSWORD_ALREADY_SET")
)

// Identity validation (1301-1399).
var (
	ErrInvalidEmail = New(connect.CodeInvalidArgument, 1301, "INVALID_EMAIL")
	ErrInvalidPhone = New(connect.CodeInvalidArgument, 1302, "INVALID_PHONE")
)

// Login method (1401-1499).
var (
	ErrLoginMethodNotAllowed = New(connect.CodeInvalidArgument, 1401, "LOGIN_METHOD_NOT_ALLOWED")
)
