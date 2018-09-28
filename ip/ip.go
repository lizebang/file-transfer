package ip

import (
	"net"
)

func IP() string {
	ips := []string(nil)

	is, err := net.Interfaces()
	if err != nil || len(is) == 0 {
		return ""
	}

	for _, i := range is {
		if len(i.HardwareAddr) == 0 {
			continue
		}
		as, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, a := range as {
			ip, ok := a.(*net.IPNet)
			if ok && !ip.IP.IsLoopback() && ip.IP.To4() != nil {
				ips = append(ips, ip.IP.String())
			}
		}
	}

	if len(ips) == 0 {
		return ""
	}
	return ips[0]
}
