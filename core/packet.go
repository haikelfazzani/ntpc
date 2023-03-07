package ntpc

import (
	"bytes"
	"encoding/binary"
)

type NtpTime struct {
	Seconds  uint32
	Fraction uint32
}

type NtpPacket struct {
	LiVnMode  uint8 // Leap indicator(2), Version number(3), Mode(3)
	Stratum   uint8 // Stratum level of the local clock
	Poll      uint8 // Maximum interval between successive messages
	Precision int8  // Representing the precision of the system clock

	RootDelay      uint32 // Round trip time to the primary reference source
	RootDispersion uint32 // Dispersion to the primary reference source
	ReferenceId    uint32 // Reference identifier

	ReferenceTime NtpTime // The time when the system clock was last set or corrected
	OriginateTime NtpTime // The time at which the request departed the client for the server
	ReceiveTime   NtpTime // The time at which the request arrived at the server
	TransmitTime  NtpTime // The time at which the reply departed the server for client
}

func (p *NtpPacket) ToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, p)
	return buf.Bytes(), err
}
