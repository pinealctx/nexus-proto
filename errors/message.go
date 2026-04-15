package errors

import "connectrpc.com/connect"

// Message error codes (6000-6999). Reference: docs/designs/d12-error-codes.md.
var (
	ErrMessageNotFound        = New(connect.CodeNotFound, 6001, "MESSAGE_NOT_FOUND")
	ErrNotMessageSender       = New(connect.CodePermissionDenied, 6002, "NOT_MESSAGE_SENDER")
	ErrRecallTimeout          = New(connect.CodeFailedPrecondition, 6003, "RECALL_TIMEOUT")
	ErrMessageAlreadyRecalled = New(connect.CodeFailedPrecondition, 6004, "MESSAGE_ALREADY_RECALLED")
	ErrInvalidMessageType     = New(connect.CodeInvalidArgument, 6005, "INVALID_MESSAGE_TYPE")
	ErrMessageTooLong         = New(connect.CodeInvalidArgument, 6006, "MESSAGE_TOO_LONG")
	ErrSendBlocked            = New(connect.CodePermissionDenied, 6007, "SEND_BLOCKED")
	ErrDuplicateMessage       = New(connect.CodeAlreadyExists, 6008, "DUPLICATE_MESSAGE")
	ErrStreamNotFound         = New(connect.CodeNotFound, 6009, "STREAM_NOT_FOUND")
	ErrStreamSeqInvalid       = New(connect.CodeInvalidArgument, 6010, "STREAM_SEQ_INVALID")
	ErrCardActionExpired      = New(connect.CodeNotFound, 6011, "CARD_ACTION_EXPIRED")
	ErrStreamAlreadyEnded     = New(connect.CodeFailedPrecondition, 6012, "STREAM_ALREADY_ENDED")
	ErrMessageTypeMismatch    = New(connect.CodeInvalidArgument, 6013, "MESSAGE_TYPE_MISMATCH")
	ErrNotCardMessage         = New(connect.CodeFailedPrecondition, 6014, "NOT_CARD_MESSAGE")
)
