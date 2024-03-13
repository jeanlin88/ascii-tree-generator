package tests

import "testing"

func Error(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("expect error get nil")
		t.FailNow()
	}
}

func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got unexpected error %q", err.Error())
		t.FailNow()
	}
}

func EqualError(t *testing.T, expect, get error) {
	t.Helper()
	if expect != get {
		t.Errorf("expect %q get %q", expect, get)
		t.FailNow()
	}
}

func EqualString(t *testing.T, expect, get string) {
	t.Helper()
	if expect != get {
		t.Errorf("expect %q get %q", expect, get)
		t.FailNow()
	}
}

func EqualBool(t *testing.T, expect, get bool) {
	t.Helper()
	if expect != get {
		t.Errorf("expect %v get %v", expect, get)
		t.FailNow()
	}
}
