package infrastructure

import (
	"reflect"
	"testing"
)

func TestNewDB(t *testing.T) {
	actual := NewDB()
	expected := &DB{}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v\n", actual, expected)
	}
}
