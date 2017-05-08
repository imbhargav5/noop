package main

import (
	"reflect"
	"testing"
)

func TestNoop(t *testing.T) {
	if reflect.TypeOf(Noop).Kind() != reflect.Func {
		t.Errorf("Expected Noop to be a function, Noop is %s", reflect.TypeOf(Noop).Kind())
	}
}
