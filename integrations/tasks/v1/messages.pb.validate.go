// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: integrations/tasks/v1/messages.proto

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

// Validate checks the field values on CreateTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTaskRequestMultiError, or nil if none found.
func (m *CreateTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetIntegrationId() != "" {

		if err := m._validateUuid(m.GetIntegrationId()); err != nil {
			err = CreateTaskRequestValidationError{
				field:  "IntegrationId",
				reason: "value must be a valid UUID",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	for idx, item := range m.GetProperties() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateTaskRequestValidationError{
						field:  fmt.Sprintf("Properties[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateTaskRequestValidationError{
						field:  fmt.Sprintf("Properties[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateTaskRequestValidationError{
					field:  fmt.Sprintf("Properties[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CreateTaskRequestMultiError(errors)
	}

	return nil
}

func (m *CreateTaskRequest) _validateUuid(uuid string) error {
	if matched := _messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// CreateTaskRequestMultiError is an error wrapping multiple validation errors
// returned by CreateTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTaskRequestMultiError) AllErrors() []error { return m }

// CreateTaskRequestValidationError is the validation error returned by
// CreateTaskRequest.Validate if the designated constraints aren't met.
type CreateTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTaskRequestValidationError) ErrorName() string {
	return "CreateTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTaskRequestValidationError{}

// Validate checks the field values on CreateTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTaskResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTaskResponseMultiError, or nil if none found.
func (m *CreateTaskResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTaskResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateTaskResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateTaskResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTaskResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateTaskResponseValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateTaskResponseValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTaskResponseValidationError{
				field:  "Task",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateTaskResponseMultiError(errors)
	}

	return nil
}

// CreateTaskResponseMultiError is an error wrapping multiple validation errors
// returned by CreateTaskResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateTaskResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTaskResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTaskResponseMultiError) AllErrors() []error { return m }

// CreateTaskResponseValidationError is the validation error returned by
// CreateTaskResponse.Validate if the designated constraints aren't met.
type CreateTaskResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTaskResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTaskResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTaskResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTaskResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTaskResponseValidationError) ErrorName() string {
	return "CreateTaskResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTaskResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTaskResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTaskResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTaskResponseValidationError{}

// Validate checks the field values on GetTaskRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTaskRequestMultiError,
// or nil if none found.
func (m *GetTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() != "" {

		if err := m._validateUuid(m.GetId()); err != nil {
			err = GetTaskRequestValidationError{
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

	if len(errors) > 0 {
		return GetTaskRequestMultiError(errors)
	}

	return nil
}

func (m *GetTaskRequest) _validateUuid(uuid string) error {
	if matched := _messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetTaskRequestMultiError is an error wrapping multiple validation errors
// returned by GetTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTaskRequestMultiError) AllErrors() []error { return m }

// GetTaskRequestValidationError is the validation error returned by
// GetTaskRequest.Validate if the designated constraints aren't met.
type GetTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskRequestValidationError) ErrorName() string { return "GetTaskRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskRequestValidationError{}

// Validate checks the field values on GetTaskResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetTaskResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTaskResponseMultiError, or nil if none found.
func (m *GetTaskResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTaskResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTaskResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTaskResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTaskResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTaskResponseValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTaskResponseValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTaskResponseValidationError{
				field:  "Task",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetProperties() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTaskResponseValidationError{
						field:  fmt.Sprintf("Properties[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTaskResponseValidationError{
						field:  fmt.Sprintf("Properties[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTaskResponseValidationError{
					field:  fmt.Sprintf("Properties[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetTaskResponseMultiError(errors)
	}

	return nil
}

// GetTaskResponseMultiError is an error wrapping multiple validation errors
// returned by GetTaskResponse.ValidateAll() if the designated constraints
// aren't met.
type GetTaskResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTaskResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTaskResponseMultiError) AllErrors() []error { return m }

// GetTaskResponseValidationError is the validation error returned by
// GetTaskResponse.Validate if the designated constraints aren't met.
type GetTaskResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskResponseValidationError) ErrorName() string { return "GetTaskResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetTaskResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskResponseValidationError{}

// Validate checks the field values on GetTaskStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTaskStatusRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTaskStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTaskStatusRequestMultiError, or nil if none found.
func (m *GetTaskStatusRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTaskStatusRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() != "" {

		if err := m._validateUuid(m.GetId()); err != nil {
			err = GetTaskStatusRequestValidationError{
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

	if len(errors) > 0 {
		return GetTaskStatusRequestMultiError(errors)
	}

	return nil
}

func (m *GetTaskStatusRequest) _validateUuid(uuid string) error {
	if matched := _messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetTaskStatusRequestMultiError is an error wrapping multiple validation
// errors returned by GetTaskStatusRequest.ValidateAll() if the designated
// constraints aren't met.
type GetTaskStatusRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTaskStatusRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTaskStatusRequestMultiError) AllErrors() []error { return m }

// GetTaskStatusRequestValidationError is the validation error returned by
// GetTaskStatusRequest.Validate if the designated constraints aren't met.
type GetTaskStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskStatusRequestValidationError) ErrorName() string {
	return "GetTaskStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTaskStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskStatusRequestValidationError{}

// Validate checks the field values on GetTaskStatusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTaskStatusResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTaskStatusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTaskStatusResponseMultiError, or nil if none found.
func (m *GetTaskStatusResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTaskStatusResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTaskStatusResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTaskStatusResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTaskStatusResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTaskStatusResponseValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTaskStatusResponseValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTaskStatusResponseValidationError{
				field:  "Task",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetStatuses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTaskStatusResponseValidationError{
						field:  fmt.Sprintf("Statuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTaskStatusResponseValidationError{
						field:  fmt.Sprintf("Statuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTaskStatusResponseValidationError{
					field:  fmt.Sprintf("Statuses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetTaskStatusResponseMultiError(errors)
	}

	return nil
}

// GetTaskStatusResponseMultiError is an error wrapping multiple validation
// errors returned by GetTaskStatusResponse.ValidateAll() if the designated
// constraints aren't met.
type GetTaskStatusResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTaskStatusResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTaskStatusResponseMultiError) AllErrors() []error { return m }

// GetTaskStatusResponseValidationError is the validation error returned by
// GetTaskStatusResponse.Validate if the designated constraints aren't met.
type GetTaskStatusResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskStatusResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskStatusResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskStatusResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskStatusResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskStatusResponseValidationError) ErrorName() string {
	return "GetTaskStatusResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetTaskStatusResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskStatusResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskStatusResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskStatusResponseValidationError{}

// Validate checks the field values on GetTaskResultRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTaskResultRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTaskResultRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTaskResultRequestMultiError, or nil if none found.
func (m *GetTaskResultRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTaskResultRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() != "" {

		if err := m._validateUuid(m.GetId()); err != nil {
			err = GetTaskResultRequestValidationError{
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

	if len(errors) > 0 {
		return GetTaskResultRequestMultiError(errors)
	}

	return nil
}

func (m *GetTaskResultRequest) _validateUuid(uuid string) error {
	if matched := _messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetTaskResultRequestMultiError is an error wrapping multiple validation
// errors returned by GetTaskResultRequest.ValidateAll() if the designated
// constraints aren't met.
type GetTaskResultRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTaskResultRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTaskResultRequestMultiError) AllErrors() []error { return m }

// GetTaskResultRequestValidationError is the validation error returned by
// GetTaskResultRequest.Validate if the designated constraints aren't met.
type GetTaskResultRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskResultRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskResultRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskResultRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskResultRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskResultRequestValidationError) ErrorName() string {
	return "GetTaskResultRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTaskResultRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskResultRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskResultRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskResultRequestValidationError{}

// Validate checks the field values on GetTaskResultResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTaskResultResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTaskResultResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTaskResultResponseMultiError, or nil if none found.
func (m *GetTaskResultResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTaskResultResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTaskResultResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTaskResultResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTaskResultResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTaskResultResponseValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTaskResultResponseValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTaskResultResponseValidationError{
				field:  "Task",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetArtifacts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTaskResultResponseValidationError{
						field:  fmt.Sprintf("Artifacts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTaskResultResponseValidationError{
						field:  fmt.Sprintf("Artifacts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTaskResultResponseValidationError{
					field:  fmt.Sprintf("Artifacts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetTaskResultResponseMultiError(errors)
	}

	return nil
}

// GetTaskResultResponseMultiError is an error wrapping multiple validation
// errors returned by GetTaskResultResponse.ValidateAll() if the designated
// constraints aren't met.
type GetTaskResultResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTaskResultResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTaskResultResponseMultiError) AllErrors() []error { return m }

// GetTaskResultResponseValidationError is the validation error returned by
// GetTaskResultResponse.Validate if the designated constraints aren't met.
type GetTaskResultResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskResultResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskResultResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskResultResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskResultResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskResultResponseValidationError) ErrorName() string {
	return "GetTaskResultResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetTaskResultResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskResultResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskResultResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskResultResponseValidationError{}
