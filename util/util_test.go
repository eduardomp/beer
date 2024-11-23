package util

import (
	"testing"
)

func TestConfig(t *testing.T) {
	want := "config.beer.yml"
	got := Config("BEER_CONFIG_PATH", want)
	if got != want {
		t.Errorf("Config() = %v, want %v", got, want)
	}
}
