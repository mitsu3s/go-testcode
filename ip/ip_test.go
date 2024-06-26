package ip

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateIPHeader(t *testing.T) {
	t.Parallel()

	testcases := []IPHeader{
		{
			Version:  4,
			IHL:      5,
			TOS:      0,
			Length:   20,
			ID:       0,
			Flags:    0,
			FragOff:  0,
			TTL:      64,
			Protocol: 6,
			Checksum: 0,
			SrcIP:    net.ParseIP("192.168.10.1"),
			DstIP:    net.ParseIP("192.168.10.10"),
		},
		// {
		// 	Version:  4,
		// 	IHL:      5,
		// 	TOS:      0,
		// 	Length:   20,
		// 	ID:       0,
		// 	Flags:    0,
		// 	FragOff:  0,
		// 	TTL:      64,
		// 	Protocol: 6,
		// 	Checksum: 0,
		// 	SrcIP:    net.ParseIP("192.168.10.1"),
		// 	DstIP:    net.ParseIP("192.168.10.10"),
		// },
	}
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

	for i, testcase := range testcases {
		assert.Equal(t, ip.Version, testcase.Version, "Testcase %d", i)
		assert.Equal(t, ip.IHL, testcase.IHL, "Testcase %d", i)
		assert.Equal(t, ip.TOS, testcase.TOS, "Testcase %d", i)
		assert.Equal(t, ip.Length, testcase.Length, "Testcase %d", i)
		assert.Equal(t, ip.ID, testcase.ID, "Testcase %d", i)
		assert.Equal(t, ip.Flags, testcase.Flags, "Testcase %d", i)
		assert.Equal(t, ip.FragOff, testcase.FragOff, "Testcase %d", i)
		assert.Equal(t, ip.TTL, testcase.TTL, "Testcase %d", i)
		assert.Equal(t, ip.Protocol, testcase.Protocol, "Testcase %d", i)
		assert.Equal(t, ip.Checksum, testcase.Checksum, "Testcase %d", i)
		assert.Equal(t, ip.SrcIP.String(), testcase.SrcIP.String(), "Testcase %d", i)
		assert.Equal(t, ip.DstIP.String(), testcase.DstIP.String(), "Testcase %d", i)
	}
}
