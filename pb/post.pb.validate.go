// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: post.proto

package pb

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

// Validate checks the field values on CreatePostRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreatePostRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Content

	// no validation rules for CategoryId

	// no validation rules for Thumbnail

	// no validation rules for Describe

	// no validation rules for Title

	return nil
}

// CreatePostRequestValidationError is the validation error returned by
// CreatePostRequest.Validate if the designated constraints aren't met.
type CreatePostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePostRequestValidationError) ErrorName() string {
	return "CreatePostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePostRequestValidationError{}

// Validate checks the field values on CreatePostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreatePostResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	return nil
}

// CreatePostResponseValidationError is the validation error returned by
// CreatePostResponse.Validate if the designated constraints aren't met.
type CreatePostResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePostResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePostResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePostResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePostResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePostResponseValidationError) ErrorName() string {
	return "CreatePostResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePostResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePostResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePostResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePostResponseValidationError{}

// Validate checks the field values on GetPostDetailRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetPostDetailRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PostId

	return nil
}

// GetPostDetailRequestValidationError is the validation error returned by
// GetPostDetailRequest.Validate if the designated constraints aren't met.
type GetPostDetailRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPostDetailRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPostDetailRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPostDetailRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPostDetailRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPostDetailRequestValidationError) ErrorName() string {
	return "GetPostDetailRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetPostDetailRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPostDetailRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPostDetailRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPostDetailRequestValidationError{}

// Validate checks the field values on GetPostDetailResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetPostDetailResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Content

	// no validation rules for CategoryId

	// no validation rules for Thumbnail

	// no validation rules for Describe

	// no validation rules for HeartCount

	// no validation rules for View

	// no validation rules for CreatedAt

	// no validation rules for Title

	// no validation rules for Id

	return nil
}

// GetPostDetailResponseValidationError is the validation error returned by
// GetPostDetailResponse.Validate if the designated constraints aren't met.
type GetPostDetailResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPostDetailResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPostDetailResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPostDetailResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPostDetailResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPostDetailResponseValidationError) ErrorName() string {
	return "GetPostDetailResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetPostDetailResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPostDetailResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPostDetailResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPostDetailResponseValidationError{}

// Validate checks the field values on GetPostsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetPostsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Offset

	// no validation rules for Limit

	return nil
}

// GetPostsRequestValidationError is the validation error returned by
// GetPostsRequest.Validate if the designated constraints aren't met.
type GetPostsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPostsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPostsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPostsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPostsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPostsRequestValidationError) ErrorName() string { return "GetPostsRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetPostsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPostsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPostsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPostsRequestValidationError{}

// Validate checks the field values on GetPostsResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetPostsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPosts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetPostsResponseValidationError{
					field:  fmt.Sprintf("Posts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	return nil
}

// GetPostsResponseValidationError is the validation error returned by
// GetPostsResponse.Validate if the designated constraints aren't met.
type GetPostsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPostsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPostsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPostsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPostsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPostsResponseValidationError) ErrorName() string { return "GetPostsResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetPostsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPostsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPostsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPostsResponseValidationError{}

// Validate checks the field values on UpdatePostRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpdatePostRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PostId

	// no validation rules for Content

	// no validation rules for CategoryId

	// no validation rules for Thumbnail

	// no validation rules for Describe

	// no validation rules for Title

	return nil
}

// UpdatePostRequestValidationError is the validation error returned by
// UpdatePostRequest.Validate if the designated constraints aren't met.
type UpdatePostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePostRequestValidationError) ErrorName() string {
	return "UpdatePostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePostRequestValidationError{}

// Validate checks the field values on UpdatePostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdatePostResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	return nil
}

// UpdatePostResponseValidationError is the validation error returned by
// UpdatePostResponse.Validate if the designated constraints aren't met.
type UpdatePostResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePostResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePostResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePostResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePostResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePostResponseValidationError) ErrorName() string {
	return "UpdatePostResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePostResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePostResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePostResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePostResponseValidationError{}

// Validate checks the field values on DeletePostRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeletePostRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PostId

	return nil
}

// DeletePostRequestValidationError is the validation error returned by
// DeletePostRequest.Validate if the designated constraints aren't met.
type DeletePostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePostRequestValidationError) ErrorName() string {
	return "DeletePostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeletePostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePostRequestValidationError{}

// Validate checks the field values on DeletePostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeletePostResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	return nil
}

// DeletePostResponseValidationError is the validation error returned by
// DeletePostResponse.Validate if the designated constraints aren't met.
type DeletePostResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePostResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePostResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePostResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePostResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePostResponseValidationError) ErrorName() string {
	return "DeletePostResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeletePostResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePostResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePostResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePostResponseValidationError{}
