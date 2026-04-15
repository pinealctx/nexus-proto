package errors

import "connectrpc.com/connect"

// Group error codes (4000-4999). Reference: docs/designs/d12-error-codes.md.
var (
	ErrGroupNotFound      = New(connect.CodeNotFound, 4001, "GROUP_NOT_FOUND")
	ErrGroupDissolved     = New(connect.CodeFailedPrecondition, 4002, "GROUP_DISSOLVED")
	ErrNotGroupMember     = New(connect.CodePermissionDenied, 4003, "NOT_GROUP_MEMBER")
	ErrNotGroupOwner      = New(connect.CodePermissionDenied, 4004, "NOT_GROUP_OWNER")
	ErrAlreadyGroupMember = New(connect.CodeAlreadyExists, 4005, "ALREADY_GROUP_MEMBER")
	ErrGroupMemberLimit   = New(connect.CodeResourceExhausted, 4006, "GROUP_MEMBER_LIMIT")
	ErrGroupAgentLimit    = New(connect.CodeResourceExhausted, 4007, "GROUP_AGENT_LIMIT")
	ErrCannotKickOwner    = New(connect.CodeInvalidArgument, 4008, "CANNOT_KICK_OWNER")
	ErrGroupNameInvalid   = New(connect.CodeInvalidArgument, 4009, "GROUP_NAME_INVALID")
	ErrUserGroupLimit     = New(connect.CodeResourceExhausted, 4010, "USER_GROUP_LIMIT")
	ErrInviteBlockedUser  = New(connect.CodePermissionDenied, 4011, "INVITE_BLOCKED_USER")
	ErrOwnerCannotLeave   = New(connect.CodePermissionDenied, 4012, "OWNER_CANNOT_LEAVE")
	ErrGroupTooFewMembers = New(connect.CodeFailedPrecondition, 4013, "GROUP_TOO_FEW_MEMBERS")
)
