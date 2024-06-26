package ip

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateIPHeader(t *testing.T) {
	t.Parallel()
	version := uint8(4)
	ihl := uint8(5)
	tos := uint8(0)
	length := uint16(20)
	id := uint16(0)
	flags := uint8(0)
	fragOff := uint16(0)
	ttl := uint8(64)
	protocol := uint8(6)
	checksum := uint16(0)
	src := net.ParseIP("192.168.10.1")
	dst := net.ParseIP("192.168.10.10")

	ip := CreateIPHeader(version, ihl, tos, length, id, flags, fragOff, ttl, protocol, checksum, src, dst)

	assert.Equal(t, ip.Version, uint8(4))
	assert.Equal(t, ip.IHL, uint8(5))
	assert.Equal(t, ip.TOS, uint8(0))
	assert.Equal(t, ip.Length, uint16(20))
	assert.Equal(t, ip.ID, uint16(0))
	assert.Equal(t, ip.Flags, uint8(0))
	assert.Equal(t, ip.FragOff, uint16(0))
	assert.Equal(t, ip.TTL, uint8(64))
	assert.Equal(t, ip.Protocol, uint8(6))
	assert.Equal(t, ip.Checksum, uint16(0))
	assert.Equal(t, ip.SrcIP.String(), "192.168.10.1")
	assert.Equal(t, ip.DstIP.String(), "192.168.10.10")
}
