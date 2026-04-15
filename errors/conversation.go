package errors

import "connectrpc.com/connect"

// Conversation error codes (5000-5999). Reference: docs/designs/d12-error-codes.md.
var (
	ErrConversationNotFound  = New(connect.CodeNotFound, 5001, "CONVERSATION_NOT_FOUND")
	ErrNotConversationMember = New(connect.CodePermissionDenied, 5002, "NOT_CONVERSATION_MEMBER")
	ErrConversationReadonly  = New(connect.CodeFailedPrecondition, 5003, "CONVERSATION_READONLY")
)
