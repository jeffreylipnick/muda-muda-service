package config

import "testing"

func TestNewConfig(t *testing.T) {
	// These tests cannot be run in parallel, due to using Setenv.
	// https://beta.pkg.go.dev/testing@master#T.Setenv
	t.Setenv("MUDA_PORT", "8080")

	conf := NewConfig()
	if conf.Port != "8080" {
		t.Error("port not properly set")
	}
}