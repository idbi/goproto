// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: integrations/owners/v1/messages.proto

package v1

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

// define the regex for a UUID once up-front
var _messages_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateOwnerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOwnerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOwnerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOwnerRequestMultiError, or nil if none found.
func (m *CreateOwnerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOwnerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() != "" {

		if err := m._validateUuid(m.GetId()); err != nil {
			err = CreateOwnerRequestValidationError{
				field:  "Id",
				reason: "value must be a valid UUID",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 255 {
		err := CreateOwnerRequestValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 255 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateOwnerRequestMultiError(errors)
	}

	return nil
}

func (m *CreateOwnerRequest) _validateUuid(uuid string) error {
	if matched := _messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// CreateOwnerRequestMultiError is an error wrapping multiple validation errors
// returned by CreateOwnerRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateOwnerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOwnerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOwnerRequestMultiError) AllErrors() []error { return m }

// CreateOwnerRequestValidationError is the validation error returned by
// CreateOwnerRequest.Validate if the designated constraints aren't met.
type CreateOwnerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOwnerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOwnerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOwnerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOwnerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOwnerRequestValidationError) ErrorName() string {
	return "CreateOwnerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOwnerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOwnerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOwnerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOwnerRequestValidationError{}

// Validate checks the field values on CreateOwnerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOwnerResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOwnerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOwnerResponseMultiError, or nil if none found.
func (m *CreateOwnerResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOwnerResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateOwnerResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateOwnerResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateOwnerResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOwner()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateOwnerResponseValidationError{
					field:  "Owner",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateOwnerResponseValidationError{
					field:  "Owner",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOwner()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateOwnerResponseValidationError{
				field:  "Owner",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateOwnerResponseMultiError(errors)
	}

	return nil
}

// CreateOwnerResponseMultiError is an error wrapping multiple validation
// errors returned by CreateOwnerResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateOwnerResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOwnerResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOwnerResponseMultiError) AllErrors() []error { return m }

// CreateOwnerResponseValidationError is the validation error returned by
// CreateOwnerResponse.Validate if the designated constraints aren't met.
type CreateOwnerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOwnerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOwnerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOwnerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOwnerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOwnerResponseValidationError) ErrorName() string {
	return "CreateOwnerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOwnerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOwnerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOwnerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOwnerResponseValidationError{}

// Validate checks the field values on GetOwnerRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetOwnerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOwnerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOwnerRequestMultiError, or nil if none found.
func (m *GetOwnerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOwnerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = GetOwnerRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetOwnerRequestMultiError(errors)
	}

	return nil
}

func (m *GetOwnerRequest) _validateUuid(uuid string) error {
	if matched := _messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetOwnerRequestMultiError is an error wrapping multiple validation errors
// returned by GetOwnerRequest.ValidateAll() if the designated constraints
// aren't met.
type GetOwnerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOwnerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOwnerRequestMultiError) AllErrors() []error { return m }

// GetOwnerRequestValidationError is the validation error returned by
// GetOwnerRequest.Validate if the designated constraints aren't met.
type GetOwnerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOwnerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOwnerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOwnerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOwnerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOwnerRequestValidationError) ErrorName() string { return "GetOwnerRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetOwnerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOwnerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOwnerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOwnerRequestValidationError{}

// Validate checks the field values on GetOwnerResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetOwnerResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOwnerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOwnerResponseMultiError, or nil if none found.
func (m *GetOwnerResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOwnerResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetOwnerResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetOwnerResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetOwnerResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOwner()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetOwnerResponseValidationError{
					field:  "Owner",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetOwnerResponseValidationError{
					field:  "Owner",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOwner()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetOwnerResponseValidationError{
				field:  "Owner",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetOwnerResponseMultiError(errors)
	}

	return nil
}

// GetOwnerResponseMultiError is an error wrapping multiple validation errors
// returned by GetOwnerResponse.ValidateAll() if the designated constraints
// aren't met.
type GetOwnerResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOwnerResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOwnerResponseMultiError) AllErrors() []error { return m }

// GetOwnerResponseValidationError is the validation error returned by
// GetOwnerResponse.Validate if the designated constraints aren't met.
type GetOwnerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOwnerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOwnerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOwnerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOwnerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOwnerResponseValidationError) ErrorName() string { return "GetOwnerResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetOwnerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOwnerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOwnerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOwnerResponseValidationError{}
