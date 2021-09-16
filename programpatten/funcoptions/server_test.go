package funcoptions

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	NewServer("localhost", 1024)
	NewServer("localhost", 2048, Protocol("udp"))
}
