package util

import (
	"strings"
	"testing"
)

func TestConfig(t *testing.T) {
	want := "config.beer.yml"
	got := GetBeerConfig()
	if !strings.HasSuffix(got, want) {
		t.Errorf("Config() = %v, want %v", got, want)
	}
}
