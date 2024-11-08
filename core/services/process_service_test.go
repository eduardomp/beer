package services

import (
	"reflect"
	"testing"
)

func TestGetProcessInfo(t *testing.T) {

	got := GetProcessInfo()

	kind := reflect.TypeOf(got).Kind()
	if kind != reflect.Slice {
		t.Errorf("GetProcessINfo() = %v, want %v", kind, reflect.Slice)
	}

}
