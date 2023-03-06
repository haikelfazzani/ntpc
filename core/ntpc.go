package ntpc

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func NtpQuery(server string, port string) (*time.Time, error) {

	addr, _ := net.ResolveUDPAddr("udp", net.JoinHostPort(server, port))
	conn, err := net.Dial("udp", addr.AddrPort().String())

	println(addr.AddrPort().String())

	if err != nil {
		fmt.Println("Error connecting to NTP server:", err)
		return nil, err
	}

	defer conn.Close()

	// send request
	ntpReq := make([]byte, 48)
	ntpReq[0] = 0x1B

	_, err = conn.Write(ntpReq)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}

	conn.SetDeadline(time.Now().Add(5 * time.Second))

	// read response
	packet := make([]byte, 48)
	_, err = conn.Read(packet)
	if err != nil {
		fmt.Println("Error receiving response:", err)
		return nil, err
	}

	fmt.Printf("\nResponse: %v", packet)
	fmt.Printf("\npacket: %v\n %v\n", packet[40:48], packet[40])
	fmt.Printf("\nuint64: %v\n %s\n", uint64(packet[40]), strconv.FormatUint(uint64(packet[40]), 10))

	// extract and convert timestamp
	second := uint64(packet[40])<<24 | uint64(packet[41])<<16 | uint64(packet[42])<<8 | uint64(packet[43])
	fraction := uint64(packet[44])<<24 | uint64(packet[45])<<16 | uint64(packet[46])<<8 | uint64(packet[47])

	nsec := (second * 1e9) + ((fraction * 1e9) >> 32)

	now := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(nsec))

	return &now, nil
}
