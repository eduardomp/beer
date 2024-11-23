package services

import (
	"beer/core/model"
	"reflect"
	"testing"
)

func TestGetProcessInfo(t *testing.T) {

	got := GetProcessInfo()

	kind := reflect.TypeOf(got).Kind()
	if kind != reflect.Slice {
		t.Errorf("GetProcessInfo() returning a non slice value %v, want %v", kind, reflect.Slice)
	}

	reflect.TypeOf(got[0]).ConvertibleTo(reflect.TypeOf(model.Process{}))
	if kind != reflect.Slice {
		t.Errorf("GetProcessINfo() returning a slice with non model.Process value %v, want %v", kind, reflect.Slice)
	}

}
