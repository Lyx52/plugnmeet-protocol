// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: plugnmeet_auth_room.proto

package plugnmeet

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GetActiveRoomInfoReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetActiveRoomInfoReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetActiveRoomInfoReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetActiveRoomInfoReqMultiError, or nil if none found.
func (m *GetActiveRoomInfoReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetActiveRoomInfoReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoomId

	if len(errors) > 0 {
		return GetActiveRoomInfoReqMultiError(errors)
	}

	return nil
}

// GetActiveRoomInfoReqMultiError is an error wrapping multiple validation
// errors returned by GetActiveRoomInfoReq.ValidateAll() if the designated
// constraints aren't met.
type GetActiveRoomInfoReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetActiveRoomInfoReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetActiveRoomInfoReqMultiError) AllErrors() []error { return m }

// GetActiveRoomInfoReqValidationError is the validation error returned by
// GetActiveRoomInfoReq.Validate if the designated constraints aren't met.
type GetActiveRoomInfoReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetActiveRoomInfoReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetActiveRoomInfoReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetActiveRoomInfoReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetActiveRoomInfoReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetActiveRoomInfoReqValidationError) ErrorName() string {
	return "GetActiveRoomInfoReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetActiveRoomInfoReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetActiveRoomInfoReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetActiveRoomInfoReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetActiveRoomInfoReqValidationError{}

// Validate checks the field values on ActiveRoomInfoRes with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ActiveRoomInfoRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ActiveRoomInfoRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ActiveRoomInfoResMultiError, or nil if none found.
func (m *ActiveRoomInfoRes) ValidateAll() error {
	return m.validate(true)
}

func (m *ActiveRoomInfoRes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Msg

	for idx, item := range m.GetParticipantsInfo() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ActiveRoomInfoResValidationError{
						field:  fmt.Sprintf("ParticipantsInfo[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ActiveRoomInfoResValidationError{
						field:  fmt.Sprintf("ParticipantsInfo[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ActiveRoomInfoResValidationError{
					field:  fmt.Sprintf("ParticipantsInfo[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.RoomInfo != nil {

		if all {
			switch v := interface{}(m.GetRoomInfo()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ActiveRoomInfoResValidationError{
						field:  "RoomInfo",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ActiveRoomInfoResValidationError{
						field:  "RoomInfo",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRoomInfo()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ActiveRoomInfoResValidationError{
					field:  "RoomInfo",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ActiveRoomInfoResMultiError(errors)
	}

	return nil
}

// ActiveRoomInfoResMultiError is an error wrapping multiple validation errors
// returned by ActiveRoomInfoRes.ValidateAll() if the designated constraints
// aren't met.
type ActiveRoomInfoResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ActiveRoomInfoResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ActiveRoomInfoResMultiError) AllErrors() []error { return m }

// ActiveRoomInfoResValidationError is the validation error returned by
// ActiveRoomInfoRes.Validate if the designated constraints aren't met.
type ActiveRoomInfoResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActiveRoomInfoResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActiveRoomInfoResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActiveRoomInfoResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActiveRoomInfoResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActiveRoomInfoResValidationError) ErrorName() string {
	return "ActiveRoomInfoResValidationError"
}

// Error satisfies the builtin error interface
func (e ActiveRoomInfoResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sActiveRoomInfoRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActiveRoomInfoResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActiveRoomInfoResValidationError{}

// Validate checks the field values on ActiveRoomInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ActiveRoomInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ActiveRoomInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ActiveRoomInfoMultiError,
// or nil if none found.
func (m *ActiveRoomInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *ActiveRoomInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoomTitle

	// no validation rules for RoomId

	// no validation rules for Sid

	// no validation rules for JoinedParticipants

	// no validation rules for IsRunning

	// no validation rules for IsRecording

	// no validation rules for IsActiveRtmp

	// no validation rules for WebhookUrl

	// no validation rules for IsBreakoutRoom

	// no validation rules for ParentRoomId

	// no validation rules for CreationTime

	// no validation rules for Metadata

	if len(errors) > 0 {
		return ActiveRoomInfoMultiError(errors)
	}

	return nil
}

// ActiveRoomInfoMultiError is an error wrapping multiple validation errors
// returned by ActiveRoomInfo.ValidateAll() if the designated constraints
// aren't met.
type ActiveRoomInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ActiveRoomInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ActiveRoomInfoMultiError) AllErrors() []error { return m }

// ActiveRoomInfoValidationError is the validation error returned by
// ActiveRoomInfo.Validate if the designated constraints aren't met.
type ActiveRoomInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActiveRoomInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActiveRoomInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActiveRoomInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActiveRoomInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActiveRoomInfoValidationError) ErrorName() string { return "ActiveRoomInfoValidationError" }

// Error satisfies the builtin error interface
func (e ActiveRoomInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sActiveRoomInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActiveRoomInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActiveRoomInfoValidationError{}

// Validate checks the field values on RoomEndReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RoomEndReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RoomEndReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RoomEndReqMultiError, or
// nil if none found.
func (m *RoomEndReq) ValidateAll() error {
	return m.validate(true)
}

func (m *RoomEndReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoomId

	if len(errors) > 0 {
		return RoomEndReqMultiError(errors)
	}

	return nil
}

// RoomEndReqMultiError is an error wrapping multiple validation errors
// returned by RoomEndReq.ValidateAll() if the designated constraints aren't met.
type RoomEndReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RoomEndReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RoomEndReqMultiError) AllErrors() []error { return m }

// RoomEndReqValidationError is the validation error returned by
// RoomEndReq.Validate if the designated constraints aren't met.
type RoomEndReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoomEndReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoomEndReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoomEndReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoomEndReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoomEndReqValidationError) ErrorName() string { return "RoomEndReqValidationError" }

// Error satisfies the builtin error interface
func (e RoomEndReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoomEndReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoomEndReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoomEndReqValidationError{}

// Validate checks the field values on RoomEndRes with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RoomEndRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RoomEndRes with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RoomEndResMultiError, or
// nil if none found.
func (m *RoomEndRes) ValidateAll() error {
	return m.validate(true)
}

func (m *RoomEndRes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Msg

	if len(errors) > 0 {
		return RoomEndResMultiError(errors)
	}

	return nil
}

// RoomEndResMultiError is an error wrapping multiple validation errors
// returned by RoomEndRes.ValidateAll() if the designated constraints aren't met.
type RoomEndResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RoomEndResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RoomEndResMultiError) AllErrors() []error { return m }

// RoomEndResValidationError is the validation error returned by
// RoomEndRes.Validate if the designated constraints aren't met.
type RoomEndResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoomEndResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoomEndResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoomEndResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoomEndResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoomEndResValidationError) ErrorName() string { return "RoomEndResValidationError" }

// Error satisfies the builtin error interface
func (e RoomEndResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoomEndRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoomEndResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoomEndResValidationError{}

// Validate checks the field values on IsRoomActiveReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IsRoomActiveReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IsRoomActiveReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IsRoomActiveReqMultiError, or nil if none found.
func (m *IsRoomActiveReq) ValidateAll() error {
	return m.validate(true)
}

func (m *IsRoomActiveReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoomId

	if len(errors) > 0 {
		return IsRoomActiveReqMultiError(errors)
	}

	return nil
}

// IsRoomActiveReqMultiError is an error wrapping multiple validation errors
// returned by IsRoomActiveReq.ValidateAll() if the designated constraints
// aren't met.
type IsRoomActiveReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IsRoomActiveReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IsRoomActiveReqMultiError) AllErrors() []error { return m }

// IsRoomActiveReqValidationError is the validation error returned by
// IsRoomActiveReq.Validate if the designated constraints aren't met.
type IsRoomActiveReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IsRoomActiveReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IsRoomActiveReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IsRoomActiveReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IsRoomActiveReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IsRoomActiveReqValidationError) ErrorName() string { return "IsRoomActiveReqValidationError" }

// Error satisfies the builtin error interface
func (e IsRoomActiveReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIsRoomActiveReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IsRoomActiveReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IsRoomActiveReqValidationError{}

// Validate checks the field values on IsRoomActiveRes with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IsRoomActiveRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IsRoomActiveRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IsRoomActiveResMultiError, or nil if none found.
func (m *IsRoomActiveRes) ValidateAll() error {
	return m.validate(true)
}

func (m *IsRoomActiveRes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Msg

	if len(errors) > 0 {
		return IsRoomActiveResMultiError(errors)
	}

	return nil
}

// IsRoomActiveResMultiError is an error wrapping multiple validation errors
// returned by IsRoomActiveRes.ValidateAll() if the designated constraints
// aren't met.
type IsRoomActiveResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IsRoomActiveResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IsRoomActiveResMultiError) AllErrors() []error { return m }

// IsRoomActiveResValidationError is the validation error returned by
// IsRoomActiveRes.Validate if the designated constraints aren't met.
type IsRoomActiveResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IsRoomActiveResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IsRoomActiveResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IsRoomActiveResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IsRoomActiveResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IsRoomActiveResValidationError) ErrorName() string { return "IsRoomActiveResValidationError" }

// Error satisfies the builtin error interface
func (e IsRoomActiveResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIsRoomActiveRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IsRoomActiveResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IsRoomActiveResValidationError{}

// Validate checks the field values on ActiveRoomWithParticipant with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ActiveRoomWithParticipant) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ActiveRoomWithParticipant with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ActiveRoomWithParticipantMultiError, or nil if none found.
func (m *ActiveRoomWithParticipant) ValidateAll() error {
	return m.validate(true)
}

func (m *ActiveRoomWithParticipant) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetParticipantsInfo() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ActiveRoomWithParticipantValidationError{
						field:  fmt.Sprintf("ParticipantsInfo[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ActiveRoomWithParticipantValidationError{
						field:  fmt.Sprintf("ParticipantsInfo[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ActiveRoomWithParticipantValidationError{
					field:  fmt.Sprintf("ParticipantsInfo[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.RoomInfo != nil {

		if all {
			switch v := interface{}(m.GetRoomInfo()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ActiveRoomWithParticipantValidationError{
						field:  "RoomInfo",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ActiveRoomWithParticipantValidationError{
						field:  "RoomInfo",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRoomInfo()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ActiveRoomWithParticipantValidationError{
					field:  "RoomInfo",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ActiveRoomWithParticipantMultiError(errors)
	}

	return nil
}

// ActiveRoomWithParticipantMultiError is an error wrapping multiple validation
// errors returned by ActiveRoomWithParticipant.ValidateAll() if the
// designated constraints aren't met.
type ActiveRoomWithParticipantMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ActiveRoomWithParticipantMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ActiveRoomWithParticipantMultiError) AllErrors() []error { return m }

// ActiveRoomWithParticipantValidationError is the validation error returned by
// ActiveRoomWithParticipant.Validate if the designated constraints aren't met.
type ActiveRoomWithParticipantValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActiveRoomWithParticipantValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActiveRoomWithParticipantValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActiveRoomWithParticipantValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActiveRoomWithParticipantValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActiveRoomWithParticipantValidationError) ErrorName() string {
	return "ActiveRoomWithParticipantValidationError"
}

// Error satisfies the builtin error interface
func (e ActiveRoomWithParticipantValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sActiveRoomWithParticipant.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActiveRoomWithParticipantValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActiveRoomWithParticipantValidationError{}

// Validate checks the field values on GetActiveRoomInfoRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetActiveRoomInfoRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetActiveRoomInfoRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetActiveRoomInfoResMultiError, or nil if none found.
func (m *GetActiveRoomInfoRes) ValidateAll() error {
	return m.validate(true)
}

func (m *GetActiveRoomInfoRes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Msg

	if all {
		switch v := interface{}(m.GetRoom()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetActiveRoomInfoResValidationError{
					field:  "Room",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetActiveRoomInfoResValidationError{
					field:  "Room",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRoom()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetActiveRoomInfoResValidationError{
				field:  "Room",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetActiveRoomInfoResMultiError(errors)
	}

	return nil
}

// GetActiveRoomInfoResMultiError is an error wrapping multiple validation
// errors returned by GetActiveRoomInfoRes.ValidateAll() if the designated
// constraints aren't met.
type GetActiveRoomInfoResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetActiveRoomInfoResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetActiveRoomInfoResMultiError) AllErrors() []error { return m }

// GetActiveRoomInfoResValidationError is the validation error returned by
// GetActiveRoomInfoRes.Validate if the designated constraints aren't met.
type GetActiveRoomInfoResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetActiveRoomInfoResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetActiveRoomInfoResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetActiveRoomInfoResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetActiveRoomInfoResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetActiveRoomInfoResValidationError) ErrorName() string {
	return "GetActiveRoomInfoResValidationError"
}

// Error satisfies the builtin error interface
func (e GetActiveRoomInfoResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetActiveRoomInfoRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetActiveRoomInfoResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetActiveRoomInfoResValidationError{}

// Validate checks the field values on GetActiveRoomsInfoRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetActiveRoomsInfoRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetActiveRoomsInfoRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetActiveRoomsInfoResMultiError, or nil if none found.
func (m *GetActiveRoomsInfoRes) ValidateAll() error {
	return m.validate(true)
}

func (m *GetActiveRoomsInfoRes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Msg

	for idx, item := range m.GetRooms() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetActiveRoomsInfoResValidationError{
						field:  fmt.Sprintf("Rooms[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetActiveRoomsInfoResValidationError{
						field:  fmt.Sprintf("Rooms[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetActiveRoomsInfoResValidationError{
					field:  fmt.Sprintf("Rooms[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetActiveRoomsInfoResMultiError(errors)
	}

	return nil
}

// GetActiveRoomsInfoResMultiError is an error wrapping multiple validation
// errors returned by GetActiveRoomsInfoRes.ValidateAll() if the designated
// constraints aren't met.
type GetActiveRoomsInfoResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetActiveRoomsInfoResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetActiveRoomsInfoResMultiError) AllErrors() []error { return m }

// GetActiveRoomsInfoResValidationError is the validation error returned by
// GetActiveRoomsInfoRes.Validate if the designated constraints aren't met.
type GetActiveRoomsInfoResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetActiveRoomsInfoResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetActiveRoomsInfoResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetActiveRoomsInfoResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetActiveRoomsInfoResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetActiveRoomsInfoResValidationError) ErrorName() string {
	return "GetActiveRoomsInfoResValidationError"
}

// Error satisfies the builtin error interface
func (e GetActiveRoomsInfoResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetActiveRoomsInfoRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetActiveRoomsInfoResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetActiveRoomsInfoResValidationError{}
