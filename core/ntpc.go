package ntpc

import (
	"net"
	"os/exec"
	"time"
)

type NTPC struct {
	Server string
	Port   string
}

func (ntpc *NTPC) Query() (*time.Time, error) {

	addr, _ := net.ResolveUDPAddr("udp", net.JoinHostPort(ntpc.Server, ntpc.Port))
	conn, err := net.Dial("udp", addr.AddrPort().String())

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	if err := conn.SetDeadline(time.Now().Add(15 * time.Second)); err != nil {
		return nil, err
	}

	ntpPacket := &NtpPacket{
		LiVnMode: 0x1b, // 0b 00 011 011
	}

	request, _ := ntpPacket.ToBytes()
	_, err = conn.Write(request)

	if err != nil {
		return nil, err
	}

	packet := make([]byte, 48)
	_, err = conn.Read(packet)
	if err != nil {
		return nil, err
	}

	// extract and convert timestamp
	second := uint64(packet[40])<<24 | uint64(packet[41])<<16 | uint64(packet[42])<<8 | uint64(packet[43])
	fraction := uint64(packet[44])<<24 | uint64(packet[45])<<16 | uint64(packet[46])<<8 | uint64(packet[47])

	nsec := (second * 1e9) + ((fraction * 1e9) >> 32)

	now := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(nsec))

	return &now, nil
}

func (ntpc *NTPC) PacketValidate(packet []byte) bool {
	return true
}

func (ntpc *NTPC) UpdateSystem(dateTime string) bool {
	_, err := exec.Command("bash", "-c", "timedatectl", "set-time", "'"+dateTime+"'").Output()
	return err == nil
}
