package services

import (
	"reflect"
	"testing"
)

func TestGetHost(t *testing.T) {

	got := GetHost()

	kind := reflect.TypeOf(got).Kind()
	if kind != reflect.Struct {
		t.Errorf("GetHost() = %v, want %v", kind, reflect.Struct)
	}

}
