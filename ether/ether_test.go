package ether

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEthernetHeader(t *testing.T) {
	t.Parallel()
	dst := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	src := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	typ := uint16(0x0800)

	eth := CreateEthernetHeader(dst, src, typ)

	assert.Equal(t, eth.DstMacAddr.String(), "00:00:00:00:00:00")
	assert.Equal(t, eth.SrcMacAddr.String(), "00:00:00:00:00:00")
	assert.Equal(t, eth.Type, uint16(0x0800))
}

func TestCreateARPHeader(t *testing.T) {
	t.Parallel()
	hwType := uint16(0x0001)
	protoType := uint16(0x0800)
	hwAddrLen := uint8(0x06)
	protoAddrLen := uint8(0x04)
	op := uint16(0x0001)
	srcMac := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	srcIP := net.ParseIP("192.168.10.1")
	dstMac := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	dstIP := net.ParseIP("192.168.10.10")

	arp := CreateARPHeader(hwType, protoType, hwAddrLen, protoAddrLen, op, srcMac, dstMac, srcIP, dstIP)

	assert.Equal(t, arp.HardwareType, uint16(0x0001))
	assert.Equal(t, arp.ProtocolType, uint16(0x0800))
	assert.Equal(t, arp.HardwareAddrLength, uint8(0x06))
	assert.Equal(t, arp.ProtocolAddrLength, uint8(0x04))
	assert.Equal(t, arp.Operation, uint16(0x0001))
	assert.Equal(t, arp.SrcMacAddr.String(), "00:00:00:00:00:00")
	assert.Equal(t, arp.SrcIPAddr.String(), "192.168.10.1")
	assert.Equal(t, arp.DstMacAddr.String(), "00:00:00:00:00:00")
	assert.Equal(t, arp.DstIPAddr.String(), "192.168.10.10")
}
