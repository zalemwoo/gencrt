package main

import (
	"flag"
	"os"

	"github.com/zalemwoo/gencrt"
)

var cn, dnss, ips string
var days int
var out string

func init() {
	flag.StringVar(&cn, "cn", "", "Common Name (required)")
	flag.IntVar(&days, "days", 365, "Days")
	flag.StringVar(&dnss, "dnss", "", "DNSNames (exp: \"example.com,*.example.com\")")
	flag.StringVar(&ips, "ips", "", "IPAddresses (exp: \"127.0.0.1,127.0.1.1\")")
	flag.StringVar(&out, "out", "", "Filename without ext. if not provided, |cn| is used instead")
}

func main() {

	flag.Parse()

	if cn == "" {
		flag.Usage()
		os.Exit(255)
	}

	if out == "" {
		out = cn
	}

	cfg := gencrt.NewConfig(cn, dnss, ips, days)

	gen, err := gencrt.NewGenerator(cfg)
	if err != nil {
		panic(err)
	}

	err = gen.WritePrivateKeyPEM(out + ".key")
	if err != nil {
		panic(err)
	}

	err = gen.WriteCertificatePEM(out + ".crt")
	if err != nil {
		panic(err)
	}
}
