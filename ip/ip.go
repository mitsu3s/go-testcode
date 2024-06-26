package ip

import "net"

type IPHeader struct {
	Version  uint8
	IHL      uint8
	TOS      uint8
	Length   uint16
	ID       uint16
	Flags    uint8
	FragOff  uint16
	TTL      uint8
	Protocol uint8
	Checksum uint16
	SrcIP    net.IP
	DstIP    net.IP
}

func CreateIPHeader(version, ihl, tos uint8, length, id uint16, flags uint8, fragOff uint16, ttl, protocol uint8, checksum uint16, src, dst net.IP) *IPHeader {
	return &IPHeader{
		Version:  version,
		IHL:      ihl,
		TOS:      tos,
		Length:   length,
		ID:       id,
		Flags:    flags,
		FragOff:  fragOff,
		TTL:      ttl,
		Protocol: protocol,
		Checksum: checksum,
		SrcIP:    src,
		DstIP:    dst,
	}
}
