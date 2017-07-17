package gencrt

import (
	"net"
	"strings"
)

type Config struct {
	CommonName  string
	Days        int
	DNSNames    []string
	IPAddresses []net.IP
}

func NewConfig(cn, dnss, ips string, days int) Config {
	cfg := Config{
		CommonName: cn,
		Days:       days,
	}

	if dnss != "" {
		dnss := strings.Split(dnss, ",")
		cfg.DNSNames = make([]string, len(dnss))
		for i, dns := range dnss {
			cfg.DNSNames[i] = strings.TrimSpace(dns)
		}
	}

	if ips != "" {
		ips := strings.Split(ips, ",")
		cfg.IPAddresses = make([]net.IP, len(ips))
		for i, ip := range ips {
			cfg.IPAddresses[i] = net.ParseIP(strings.TrimSpace(ip))
		}
	}

	return cfg
}
