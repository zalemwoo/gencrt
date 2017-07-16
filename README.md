gencrt
===

a self signed x509 certificate generator.
just for study purpose.

Usage:
```bash
Usage of ./gencrt:
  -cn string
    Common Name (required)
  -days int
    Days (default 365)
  -dnss string
    DNSNames (exp: "www.example.com,*.example.com")
  -ips string
    IPAddresses (exp: "127.0.0.1,127.0.1.1")
  -out string
    Filename without ext. if not provided, |cn| is used instead
```
