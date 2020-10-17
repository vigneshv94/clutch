// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: chaos/serverexperimentation/v1/serverexperimentation.proto

package serverexperimentationv1

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _serverexperimentation_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on TestConfig with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *TestConfig) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetClusterPair()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TestConfigValidationError{
				field:  "ClusterPair",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for FaultInjectionType

	switch m.Fault.(type) {

	case *TestConfig_Abort:

		if v, ok := interface{}(m.GetAbort()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TestConfigValidationError{
					field:  "Abort",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *TestConfig_Latency:

		if v, ok := interface{}(m.GetLatency()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TestConfigValidationError{
					field:  "Latency",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// TestConfigValidationError is the validation error returned by
// TestConfig.Validate if the designated constraints aren't met.
type TestConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TestConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TestConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TestConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TestConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TestConfigValidationError) ErrorName() string { return "TestConfigValidationError" }

// Error satisfies the builtin error interface
func (e TestConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTestConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TestConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TestConfigValidationError{}

// Validate checks the field values on ClusterPairTarget with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ClusterPairTarget) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetDownstreamCluster()) < 1 {
		return ClusterPairTargetValidationError{
			field:  "DownstreamCluster",
			reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetUpstreamCluster()) < 1 {
		return ClusterPairTargetValidationError{
			field:  "UpstreamCluster",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// ClusterPairTargetValidationError is the validation error returned by
// ClusterPairTarget.Validate if the designated constraints aren't met.
type ClusterPairTargetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterPairTargetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterPairTargetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterPairTargetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterPairTargetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterPairTargetValidationError) ErrorName() string {
	return "ClusterPairTargetValidationError"
}

// Error satisfies the builtin error interface
func (e ClusterPairTargetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterPairTarget.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterPairTargetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterPairTargetValidationError{}

// Validate checks the field values on AbortFaultConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AbortFaultConfig) Validate() error {
	if m == nil {
		return nil
	}

	if val := m.GetPercent(); val <= 0 || val > 100 {
		return AbortFaultConfigValidationError{
			field:  "Percent",
			reason: "value must be inside range (0, 100]",
		}
	}

	if val := m.GetHttpStatus(); val <= 99 || val >= 600 {
		return AbortFaultConfigValidationError{
			field:  "HttpStatus",
			reason: "value must be inside range (99, 600)",
		}
	}

	return nil
}

// AbortFaultConfigValidationError is the validation error returned by
// AbortFaultConfig.Validate if the designated constraints aren't met.
type AbortFaultConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AbortFaultConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AbortFaultConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AbortFaultConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AbortFaultConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AbortFaultConfigValidationError) ErrorName() string { return "AbortFaultConfigValidationError" }

// Error satisfies the builtin error interface
func (e AbortFaultConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAbortFaultConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AbortFaultConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AbortFaultConfigValidationError{}

// Validate checks the field values on LatencyFaultConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LatencyFaultConfig) Validate() error {
	if m == nil {
		return nil
	}

	if val := m.GetPercent(); val <= 0 || val > 100 {
		return LatencyFaultConfigValidationError{
			field:  "Percent",
			reason: "value must be inside range (0, 100]",
		}
	}

	if m.GetDurationMs() <= 0 {
		return LatencyFaultConfigValidationError{
			field:  "DurationMs",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// LatencyFaultConfigValidationError is the validation error returned by
// LatencyFaultConfig.Validate if the designated constraints aren't met.
type LatencyFaultConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LatencyFaultConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LatencyFaultConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LatencyFaultConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LatencyFaultConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LatencyFaultConfigValidationError) ErrorName() string {
	return "LatencyFaultConfigValidationError"
}

// Error satisfies the builtin error interface
func (e LatencyFaultConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLatencyFaultConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LatencyFaultConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LatencyFaultConfigValidationError{}
