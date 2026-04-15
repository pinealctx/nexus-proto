package errors

import "connectrpc.com/connect"

// User error codes (2000-2999). Reference: docs/designs/d12-error-codes.md.
var (
	ErrUserNotFound           = New(connect.CodeNotFound, 2001, "USER_NOT_FOUND")
	ErrUsernameTaken          = New(connect.CodeAlreadyExists, 2002, "USERNAME_TAKEN")
	ErrUsernameAlreadyChanged = New(connect.CodeFailedPrecondition, 2003, "USERNAME_ALREADY_CHANGED")
	ErrInvalidUsernameFormat  = New(connect.CodeInvalidArgument, 2004, "INVALID_USERNAME_FORMAT")
	ErrPhoneAlreadyBound      = New(connect.CodeAlreadyExists, 2005, "PHONE_ALREADY_BOUND")
	ErrEmailAlreadyBound      = New(connect.CodeAlreadyExists, 2006, "EMAIL_ALREADY_BOUND")
	ErrNicknameTooLong        = New(connect.CodeInvalidArgument, 2007, "NICKNAME_TOO_LONG")
	ErrSignatureTooLong       = New(connect.CodeInvalidArgument, 2008, "SIGNATURE_TOO_LONG")
	ErrDeviceNotFound         = New(connect.CodeNotFound, 2009, "DEVICE_NOT_FOUND")
	ErrCannotRemoveCurrentDev = New(connect.CodeInvalidArgument, 2010, "CANNOT_REMOVE_CURRENT_DEVICE")
	ErrUsernameReserved       = New(connect.CodeInvalidArgument, 2011, "USERNAME_RESERVED")
)
