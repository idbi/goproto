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
