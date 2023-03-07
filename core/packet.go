package ntpc

import (
	"bytes"
	"encoding/binary"
)

type NtpPacket struct {
	LiVnMode       uint8  // Leap indicator, version number, and mode
	Stratum        uint8  // Stratum level of the local clock
	Poll           uint8  // Maximum interval between successive messages
	Precision      int8   // Precision of the local clock
	RootDelay      uint32 // Round trip time to the primary reference source
	RootDispersion uint32 // Dispersion to the primary reference source
	ReferenceId    uint32 // Reference identifier
	ReferenceTime  uint64 // The time when the system clock was last set or corrected
	OriginateTime  uint64 // The time at which the request departed the client for the server
	ReceiveTime    uint64 // The time at which the request arrived at the server
	TransmitTime   uint64 // The time at which the reply departed the server for client
}

func (p *NtpPacket) ToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, p)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
