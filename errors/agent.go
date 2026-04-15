package errors

import "connectrpc.com/connect"

// Agent error codes (7000-7999). Reference: docs/designs/d12-error-codes.md.
var (
	ErrAgentNotFound             = New(connect.CodeNotFound, 7001, "AGENT_NOT_FOUND")
	ErrAgentDeleted              = New(connect.CodeFailedPrecondition, 7003, "AGENT_DELETED")
	ErrAgentUsernameTaken        = New(connect.CodeAlreadyExists, 7004, "AGENT_USERNAME_TAKEN")
	ErrNotAgentCreator           = New(connect.CodePermissionDenied, 7005, "NOT_AGENT_CREATOR")
	ErrWebhookVerificationFailed = New(connect.CodeFailedPrecondition, 7006, "WEBHOOK_VERIFICATION_FAILED")
	ErrWebhookNotConfigured      = New(connect.CodeFailedPrecondition, 7007, "WEBHOOK_NOT_CONFIGURED")
	ErrIPNotAllowed              = New(connect.CodePermissionDenied, 7008, "IP_NOT_ALLOWED")
	ErrCommandsLimitExceeded     = New(connect.CodeResourceExhausted, 7009, "COMMANDS_LIMIT_EXCEEDED")
	ErrAgentrootAlreadyExists    = New(connect.CodeAlreadyExists, 7010, "AGENTROOT_ALREADY_EXISTS")
	ErrInvalidDeliveryConfig     = New(connect.CodeInvalidArgument, 7011, "INVALID_DELIVERY_CONFIG")
	ErrDeliveryModeMismatch      = New(connect.CodeFailedPrecondition, 7012, "DELIVERY_MODE_MISMATCH")
	ErrMiniAppNotEnabled         = New(connect.CodeFailedPrecondition, 7013, "MINI_APP_NOT_ENABLED")
	ErrInvalidMiniAppURL         = New(connect.CodeInvalidArgument, 7014, "INVALID_MINI_APP_URL")
	ErrMiniAppAccessDenied       = New(connect.CodePermissionDenied, 7015, "MINI_APP_ACCESS_DENIED")
)
