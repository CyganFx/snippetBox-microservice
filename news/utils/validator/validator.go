package validator

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Validator struct {
	Errors Errors
}

func New() *Validator {
	return &Validator{
		Errors{},
	}
}

func (v *Validator) Required(fields ...string) {
	for _, field := range fields {
		if strings.TrimSpace(field) == "" {
			v.Errors.Add("Field cannot be blank")
		}
	}
}

func (v *Validator) MinLength(field string, d int) {
	if field == "" {
		return
	}
	if utf8.RuneCountInString(field) < d {
		v.Errors.Add(fmt.Sprintf("Field is too short (minimum is %d characters)", d))
	}
}

func (v *Validator) MatchesPattern(field string, pattern *regexp.Regexp) {
	if field == "" {
		return
	}
	if !pattern.MatchString(field) {
		v.Errors.Add("Invalid Field")
	}
}

func (v *Validator) MaxLength(field string, d int) {
	if field == "" {
		return
	}
	if utf8.RuneCountInString(field) > d {
		v.Errors.Add(fmt.Sprintf("This field is too long (maximum is %d characters)", d))
	}
}

func (v *Validator) PermittedValues(field string, opts ...string) {
	if field == "" {
		return
	}
	for _, opt := range opts {
		if field == opt {
			return
		}
	}
	v.Errors.Add("Invalid Field")
}

func (v *Validator) Valid() bool {
	return len(v.Errors.Errors) == 0
}
