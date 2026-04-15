package errors

import "connectrpc.com/connect"

// Contact error codes (3000-3999). Reference: docs/designs/d12-error-codes.md.
var (
	ErrAlreadyContacts       = New(connect.CodeAlreadyExists, 3001, "ALREADY_CONTACTS")
	ErrFriendRequestExists   = New(connect.CodeAlreadyExists, 3002, "FRIEND_REQUEST_EXISTS")
	ErrFriendRequestNotFound = New(connect.CodeNotFound, 3003, "FRIEND_REQUEST_NOT_FOUND")
	ErrFriendRequestHandled  = New(connect.CodeFailedPrecondition, 3004, "FRIEND_REQUEST_HANDLED")
	ErrCannotAddSelf         = New(connect.CodeInvalidArgument, 3005, "CANNOT_ADD_SELF")
	ErrUserBlocked           = New(connect.CodePermissionDenied, 3006, "USER_BLOCKED")
	ErrNotContacts           = New(connect.CodeFailedPrecondition, 3007, "NOT_CONTACTS")
	ErrAlreadyBlocked        = New(connect.CodeAlreadyExists, 3008, "ALREADY_BLOCKED")
	ErrNotBlocked            = New(connect.CodeFailedPrecondition, 3009, "NOT_BLOCKED")
	ErrAgentPrivate          = New(connect.CodePermissionDenied, 3010, "AGENT_PRIVATE")
	ErrUsesFriendRequest     = New(connect.CodeFailedPrecondition, 3011, "USE_SEND_FRIEND_REQUEST")
	ErrUsesAddContact        = New(connect.CodeFailedPrecondition, 3012, "USE_ADD_CONTACT")
)
