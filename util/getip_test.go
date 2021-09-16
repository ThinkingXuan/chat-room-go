package util

import "testing"

func TestGetLocalIP(t *testing.T) {
	t.Log(GetLocalIP())
}

func TestGetLocalShortIP(t *testing.T) {
	t.Log(GetLocalShortIP())
}
