package errors

import "connectrpc.com/connect"

// Media error codes (8000-8999). Reference: docs/designs/d12-error-codes.md.
var (
	ErrFileTooLarge          = New(connect.CodeResourceExhausted, 8001, "FILE_TOO_LARGE")
	ErrUnsupportedFileType   = New(connect.CodeInvalidArgument, 8002, "UNSUPPORTED_FILE_TYPE")
	ErrUploadSessionExpired  = New(connect.CodeNotFound, 8003, "UPLOAD_SESSION_EXPIRED")
	ErrUploadSessionNotFound = New(connect.CodeNotFound, 8004, "UPLOAD_SESSION_NOT_FOUND")
	ErrChunkOffsetInvalid    = New(connect.CodeInvalidArgument, 8005, "CHUNK_OFFSET_INVALID")
	ErrFileIncomplete        = New(connect.CodeFailedPrecondition, 8006, "FILE_INCOMPLETE")
	ErrInvalidPurpose        = New(connect.CodeInvalidArgument, 8007, "INVALID_PURPOSE")
	ErrFileNotFound          = New(connect.CodeNotFound, 8008, "FILE_NOT_FOUND")
	ErrChunkTooSmall         = New(connect.CodeInvalidArgument, 8009, "CHUNK_TOO_SMALL")
)
