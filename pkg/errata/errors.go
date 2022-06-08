// Package errata is auto-generated by errata
// Version: c1554e37487609ac8872a0be375a9bc9
package errata

import (
	"crypto/sha1"
	"fmt"
	"time"
)

type Error struct {
	Code       string
	Message    string
	Cause      string
	Categories []string
	Args       []interface{}
	Labels     map[string]string
	Guide      string

	uuid    string
	wrapped error
}

func (e Error) Unwrap() error {
	return e.wrapped
}

func (e Error) UUID() string {
	if e.uuid == "" {
		e.uuid = generateReference(e.Code)
	}
	return e.uuid
}

func (e Error) Error() string {
	message := fmt.Sprintf("[err-mimir-%s] %s. For more details, see %s", e.Code, e.Message, e.HelpURL())
	return fmt.Sprintf(message, e.Args...)
}

func (e Error) HelpURL() string {
	return fmt.Sprintf("%s%s", "https://grafana.com/docs/mimir/latest/errata/", e.Code)
}

const (
	ErrLabelNameTooLong  = "label-name-too-long"
	ErrLabelValueTooLong = "label-value-too-long"
)

var list = map[string]Error{
	ErrLabelNameTooLong: {
		Code:       ErrLabelNameTooLong,
		Message:    `Received a series whose label name length exceeds the limit, label: '%.200s' series: '%.200s'`,
		Cause:      ``,
		Categories: []string{"validation", "label"},
		Labels: map[string]string{
			"flag":               "validation.max-length-label-name",
			"http_response_code": "400",
			"level":              "warning",
		},
		Guide: `file://guides/label-name-too-long.md`,
	},

	ErrLabelValueTooLong: {
		Code:       ErrLabelValueTooLong,
		Message:    `received a series whose label value length exceeds the limit, value: '%.200s' (truncated) series: '%.200s'`,
		Cause:      ``,
		Categories: []string{"validation", "label"},
		Labels: map[string]string{
			"flag":               "validation.max-length-label-value",
			"http_response_code": "400",
			"level":              "warning",
		},
		Guide: `file://guides/label-name-too-long.md`,
	},
}

func NewFromCode(code string, wrapped error) Error {
	err := list[code]
	err.wrapped = wrapped
	return err
}

func generateReference(code string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(code+time.Now().Format(time.RFC3339Nano))))
}

func NewLabelNameTooLongErr(wrapped error, label string, series string) Error {
	err := NewFromCode(ErrLabelNameTooLong, wrapped)
	err.Args = []interface{}{label, series}
	return err
}

func NewLabelValueTooLongErr(wrapped error, label string, series string) Error {
	err := NewFromCode(ErrLabelValueTooLong, wrapped)
	err.Args = []interface{}{label, series}
	return err
}
