//go:build !linux && !windows && !darwin
// +build !linux,!windows,!darwin

package tuntap

import (
	"net"
	"testing"
)

func setupIfce(t *testing.T, ipNet net.IPNet, dev string) {
	t.Fatal("unsupported platform")
}
