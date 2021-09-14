// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: impl/v1/impl.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on DecodedZookie with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DecodedZookie) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Version

	switch m.VersionOneof.(type) {

	case *DecodedZookie_V1:

		if v, ok := interface{}(m.GetV1()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DecodedZookieValidationError{
					field:  "V1",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DecodedZookie_V2:

		if v, ok := interface{}(m.GetV2()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DecodedZookieValidationError{
					field:  "V2",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// DecodedZookieValidationError is the validation error returned by
// DecodedZookie.Validate if the designated constraints aren't met.
type DecodedZookieValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecodedZookieValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecodedZookieValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecodedZookieValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecodedZookieValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecodedZookieValidationError) ErrorName() string { return "DecodedZookieValidationError" }

// Error satisfies the builtin error interface
func (e DecodedZookieValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecodedZookie.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecodedZookieValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecodedZookieValidationError{}

// Validate checks the field values on DecodedZedToken with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DecodedZedToken) Validate() error {
	if m == nil {
		return nil
	}

	switch m.VersionOneof.(type) {

	case *DecodedZedToken_V1:

		if v, ok := interface{}(m.GetV1()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DecodedZedTokenValidationError{
					field:  "V1",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// DecodedZedTokenValidationError is the validation error returned by
// DecodedZedToken.Validate if the designated constraints aren't met.
type DecodedZedTokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecodedZedTokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecodedZedTokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecodedZedTokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecodedZedTokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecodedZedTokenValidationError) ErrorName() string { return "DecodedZedTokenValidationError" }

// Error satisfies the builtin error interface
func (e DecodedZedTokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecodedZedToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecodedZedTokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecodedZedTokenValidationError{}

// Validate checks the field values on DocComment with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *DocComment) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Comment

	return nil
}

// DocCommentValidationError is the validation error returned by
// DocComment.Validate if the designated constraints aren't met.
type DocCommentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DocCommentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DocCommentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DocCommentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DocCommentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DocCommentValidationError) ErrorName() string { return "DocCommentValidationError" }

// Error satisfies the builtin error interface
func (e DocCommentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDocComment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DocCommentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DocCommentValidationError{}

// Validate checks the field values on RelationMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RelationMetadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Kind

	return nil
}

// RelationMetadataValidationError is the validation error returned by
// RelationMetadata.Validate if the designated constraints aren't met.
type RelationMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RelationMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RelationMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RelationMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RelationMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RelationMetadataValidationError) ErrorName() string { return "RelationMetadataValidationError" }

// Error satisfies the builtin error interface
func (e RelationMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRelationMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RelationMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RelationMetadataValidationError{}

// Validate checks the field values on DecodedZookie_V1Zookie with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DecodedZookie_V1Zookie) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Revision

	return nil
}

// DecodedZookie_V1ZookieValidationError is the validation error returned by
// DecodedZookie_V1Zookie.Validate if the designated constraints aren't met.
type DecodedZookie_V1ZookieValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecodedZookie_V1ZookieValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecodedZookie_V1ZookieValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecodedZookie_V1ZookieValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecodedZookie_V1ZookieValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecodedZookie_V1ZookieValidationError) ErrorName() string {
	return "DecodedZookie_V1ZookieValidationError"
}

// Error satisfies the builtin error interface
func (e DecodedZookie_V1ZookieValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecodedZookie_V1Zookie.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecodedZookie_V1ZookieValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecodedZookie_V1ZookieValidationError{}

// Validate checks the field values on DecodedZookie_V2Zookie with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DecodedZookie_V2Zookie) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Revision

	return nil
}

// DecodedZookie_V2ZookieValidationError is the validation error returned by
// DecodedZookie_V2Zookie.Validate if the designated constraints aren't met.
type DecodedZookie_V2ZookieValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecodedZookie_V2ZookieValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecodedZookie_V2ZookieValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecodedZookie_V2ZookieValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecodedZookie_V2ZookieValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecodedZookie_V2ZookieValidationError) ErrorName() string {
	return "DecodedZookie_V2ZookieValidationError"
}

// Error satisfies the builtin error interface
func (e DecodedZookie_V2ZookieValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecodedZookie_V2Zookie.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecodedZookie_V2ZookieValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecodedZookie_V2ZookieValidationError{}

// Validate checks the field values on DecodedZedToken_V1ZedToken with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DecodedZedToken_V1ZedToken) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Revision

	return nil
}

// DecodedZedToken_V1ZedTokenValidationError is the validation error returned
// by DecodedZedToken_V1ZedToken.Validate if the designated constraints aren't met.
type DecodedZedToken_V1ZedTokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecodedZedToken_V1ZedTokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecodedZedToken_V1ZedTokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecodedZedToken_V1ZedTokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecodedZedToken_V1ZedTokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecodedZedToken_V1ZedTokenValidationError) ErrorName() string {
	return "DecodedZedToken_V1ZedTokenValidationError"
}

// Error satisfies the builtin error interface
func (e DecodedZedToken_V1ZedTokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecodedZedToken_V1ZedToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecodedZedToken_V1ZedTokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecodedZedToken_V1ZedTokenValidationError{}
