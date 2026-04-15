// Package errors provides the Nexus business error system.
//
// It is the single source of truth for all error codes, error names,
// and their Connect RPC status code mappings. Both server-side (nexus-ai)
// and client-side (nexus-x, nexus-desktop, nexus-openclaw) import this
// package instead of defining their own error constants.
//
// Server usage:
//
//	return nxerr.ErrMessageNotFound
//	// handler layer:
//	return nil, nxerr.ToConnect(err)
//
// Client usage:
//
//	if nxerr.Is(err, nxerr.ErrRecallTimeout) { ... }
//	if nxerr.IsCode(err, nxerr.ErrCardActionExpired.Code) { ... }
package errors

import (
	"errors"

	"connectrpc.com/connect"
	sharedv1 "github.com/pinealctx/nexus-proto/gen/go/shared/v1"
	"google.golang.org/protobuf/proto"
)

// Error is a Nexus business error. It binds error_code, error_name,
// and Connect status code together as a single source of truth.
type Error struct {
	ConnectCode connect.Code
	Code        int32
	Name        string
	Metadata    map[string]string
}

func (e *Error) Error() string { return e.Name }

// New creates a business error definition.
func New(connectCode connect.Code, code int32, name string) *Error {
	return &Error{ConnectCode: connectCode, Code: code, Name: name}
}

// WithMeta returns a shallow copy of the error with the given metadata
// key-value pair added. Safe to call on package-level sentinel errors
// because it never mutates the receiver.
func (e *Error) WithMeta(key, value string) *Error {
	cp := *e
	cp.Metadata = make(map[string]string, len(e.Metadata)+1)
	for k, v := range e.Metadata {
		cp.Metadata[k] = v
	}
	cp.Metadata[key] = value
	return &cp
}

// --- Server side: construct Connect errors ---

// ToConnect converts any error to a *connect.Error.
//   - *Error -> connect.Error with ErrorDetail in details.
//   - other  -> connect.Error with CodeInternal, no details.
func ToConnect(err error) *connect.Error {
	if err == nil {
		return nil
	}
	var xe *Error
	if !errors.As(err, &xe) {
		return connect.NewError(connect.CodeInternal, err)
	}
	ce := connect.NewError(xe.ConnectCode, err)
	detail, detailErr := connect.NewErrorDetail(errorToProto(xe))
	if detailErr == nil {
		ce.AddDetail(detail)
	}
	return ce
}

func errorToProto(e *Error) proto.Message {
	return &sharedv1.ErrorDetail{
		ErrorCode: e.Code,
		ErrorName: e.Name,
		Metadata:  e.Metadata,
	}
}

// --- Client side: parse Connect errors ---

// Detail holds the parsed business error from a Connect error response.
type Detail struct {
	ConnectCode connect.Code
	Code        int32
	Name        string
	Metadata    map[string]string
}

// Parse extracts the business Detail from a Connect error.
// Returns nil if the error is not a Connect error or has no ErrorDetail.
func Parse(err error) *Detail {
	if err == nil {
		return nil
	}
	var ce *connect.Error
	if !errors.As(err, &ce) {
		return nil
	}
	for _, d := range ce.Details() {
		msg, unmarshalErr := d.Value()
		if unmarshalErr != nil {
			continue
		}
		if ed, ok := msg.(*sharedv1.ErrorDetail); ok {
			return &Detail{
				ConnectCode: ce.Code(),
				Code:        ed.GetErrorCode(),
				Name:        ed.GetErrorName(),
				Metadata:    ed.GetMetadata(),
			}
		}
	}
	return nil
}

// IsCode checks if an error contains a specific business error code.
func IsCode(err error, code int32) bool {
	ed := Parse(err)
	return ed != nil && ed.Code == code
}

// Is checks if an error matches a specific Error definition by code.
func Is(err error, target *Error) bool {
	return IsCode(err, target.Code)
}
