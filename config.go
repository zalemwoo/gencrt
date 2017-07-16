package gencrt

import "net"

type Config struct {
	CommonName  string
	Days        int
	DNSNames    []string
	IPAddresses []net.IP
}
