package checkport

import (
	"fmt"
	"net"
	"time"
)

// CheckPort checks if a port is open
func CheckPort(protocol string, host string, port string, timeout time.Duration) error {
	if !(protocol == "udp" || protocol == "tcp") {
		return fmt.Errorf("protocol must be either udp or tcp")
	}

	conn, err := net.DialTimeout(protocol, net.JoinHostPort(host, port), timeout)
	if err != nil {
		return err
	}
	if conn != nil {
		defer conn.Close()
		return nil
	}
	return nil
}
