package ntpc

import (
	"encoding/binary"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"time"
)

type NewClient struct {
	Server  string
	Port    string
	Timeout time.Duration
}

func (ntpc *NewClient) Query() (*time.Time, error) {

	addr, _ := net.ResolveUDPAddr("udp", net.JoinHostPort(ntpc.Server, ntpc.Port))
	conn, err := net.Dial("udp", addr.AddrPort().String())

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	if err := conn.SetDeadline(time.Now().Add(ntpc.Timeout * time.Second)); err != nil {
		return nil, err
	}

	ntpPacket := &Packet{
		LiVnMode: 0x1b, // 0b 00 011 011
	}

	request, _ := ntpPacket.ToBytes()
	_, err = conn.Write(request)

	if err != nil {
		return nil, err
	}

	response := &Packet{}
	if err := binary.Read(conn, binary.BigEndian, response); err != nil {
		return nil, err
	}

	return response.ReceiveTime.UTC(), nil
}

func (t *NtpTime) UTC() *time.Time {
	// On POSIX-compliant OS, time is expressed using the Unix time epoch (or secs since year 1970).
	// NTP seconds are counted since 1900 and therefore must be corrected with an epoch offset to convert NTP seconds
	// to Unix time by removing 70 yrs of seconds (1970-1900) or 2208988800 seconds.
	// secs := float64(t.Seconds) - ntpEpochOffset
	// nanos := (int64(t.Fraction) * 1e9) >> 32 // convert fractional to nanos

	nsec := uint64(t.Seconds)*1e9 + (uint64(t.Fraction) * 1e9 >> 32)
	dateTime := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(nsec))

	return &dateTime
}

func (ntpc *NewClient) UpdateSystem(dateTime string) bool {

	var err error

	switch os := runtime.GOOS; os {

	case "linux":
		_, err = exec.Command("bash", "-c", "date", "-s", "'"+dateTime+"'").Output()

	case "darwin":
		_, err = exec.Command("bash", "-c", "sudo", "date", "-s", "'"+dateTime+"'").Output()
	}

	if err != nil {
		fmt.Printf("ERR: Please start the program as an administrator! %v", err)
		return false
	}
	return true
}
