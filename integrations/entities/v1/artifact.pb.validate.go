// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: integrations/entities/v1/artifact.proto

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

// Validate checks the field values on Artifact with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Artifact) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Artifact with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ArtifactMultiError, or nil
// if none found.
func (m *Artifact) ValidateAll() error {
	return m.validate(true)
}

func (m *Artifact) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Path

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ArtifactValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ArtifactValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ArtifactValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ArtifactMultiError(errors)
	}

	return nil
}

// ArtifactMultiError is an error wrapping multiple validation errors returned
// by Artifact.ValidateAll() if the designated constraints aren't met.
type ArtifactMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArtifactMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArtifactMultiError) AllErrors() []error { return m }

// ArtifactValidationError is the validation error returned by
// Artifact.Validate if the designated constraints aren't met.
type ArtifactValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArtifactValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArtifactValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArtifactValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArtifactValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArtifactValidationError) ErrorName() string { return "ArtifactValidationError" }

// Error satisfies the builtin error interface
func (e ArtifactValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArtifact.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArtifactValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArtifactValidationError{}
