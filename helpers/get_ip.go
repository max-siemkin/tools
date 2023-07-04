package tools

import (
	"net"
)

func GetMyIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		ifaces, err := net.Interfaces()
		if err != nil {
			return nil, err
		}
		for _, i := range ifaces {
			addrs, err := i.Addrs()
			if err != nil {
				return nil, err
			}
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				if ip != nil {
					return ip, nil
				}
			}
		}
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}
