package main

import (
	"reflect"
	"testing"
)

func TestIPRateLimit(t *testing.T) {
	f := IPRateLimit()
	expect := "echo.MiddlewareFunc"
	actual := reflect.TypeOf(f).String()
	if actual != expect {
		t.Errorf("this mismatch, expect: %v, got: %v", expect, actual)
	}
}
