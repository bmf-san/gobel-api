package infrastructure

import (
	"errors"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestNewLogger(t *testing.T) {
	expected := &Logger{
		errorLogFile:  "./log/error.log",
		accessLogFile: "./log/access.log",
	}
	actual := NewLogger()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v", actual, expected)
	}
}

func TestLogError(t *testing.T) {
	logger := &Logger{
		errorLogFile:  "./fixtures/logger/dummy_error.log",
		accessLogFile: "./fixtures/logger/dummy_access.log",
	}

	e := errors.New("error log")
	logger.LogError(e)

	if err := os.Remove("./fixtures/logger/dummy_error.log"); err != nil {
		t.Errorf("%v", err)
	}
}

func TestLogAccess(t *testing.T) {
	logger := &Logger{
		errorLogFile:  "./fixtures/logger/dummy_error.log",
		accessLogFile: "./fixtures/logger/dummy_access.log",
	}

	r, err := http.NewRequest("GET", "/access/", nil)
	if err != nil {
		t.Fatal(err)
	}

	logger.LogAccess(r)

	if err := os.Remove("./fixtures/logger/dummy_access.log"); err != nil {
		t.Errorf("%v", err)
	}
}
