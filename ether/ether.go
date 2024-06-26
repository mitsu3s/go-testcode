package ether

import "net"

type EthernetHeader struct {
	DstMacAddr net.HardwareAddr
	SrcMacAddr net.HardwareAddr
	Type       uint16
}

type ARPHeader struct {
	HardwareType       uint16
	ProtocolType       uint16
	HardwareAddrLength uint8
	ProtocolAddrLength uint8
	Operation          uint16
	SrcMacAddr         net.HardwareAddr
	SrcIPAddr          net.IP
	DstMacAddr         net.HardwareAddr
	DstIPAddr          net.IP
}

func CreateEthernetHeader(dst, src net.HardwareAddr, typ uint16) *EthernetHeader {
	return &EthernetHeader{
		DstMacAddr: dst,
		SrcMacAddr: src,
		Type:       typ,
	}
}

func CreateARPHeader(hwType, protoType uint16, hwAddrLen, protoAddrLen uint8, op uint16, srcMac, dstMac net.HardwareAddr, srcIPAddr, dstIPAddr net.IP) *ARPHeader {
	return &ARPHeader{
		HardwareType:       hwType,
		ProtocolType:       protoType,
		HardwareAddrLength: hwAddrLen,
		ProtocolAddrLength: protoAddrLen,
		Operation:          op,
		SrcMacAddr:         srcMac,
		SrcIPAddr:          srcIPAddr,
		DstMacAddr:         dstMac,
		DstIPAddr:          dstIPAddr,
	}
}
