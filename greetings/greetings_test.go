package greetings

import (
    "testing"
    "regexp"
)

func TestHelloName(t *testing.T) {
    name := "nick"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("nick")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}
