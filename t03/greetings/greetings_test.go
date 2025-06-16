package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	var name = "aoaochan"
	var want = regexp.MustCompile(`\b` + name + `b`)
	var msg, err = Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("aoaochan") = %q %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	var msg, err = Hello("")
	if msg != "" || err != nil {
		t.Errorf(`Hello("") = %q %v, want "" error`, msg, err)
	}
}
