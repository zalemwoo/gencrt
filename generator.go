package gencrt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"time"
)

const (
	RSA_BITLEN               = 2048
	PEM_RSA_PRIVATE_KEY_TYPE = "RSA PRIVATE KEY"
	PEM_CERTIFICATE_TYPE     = "CERTIFICATE"
)

// NewGenerator return a new generator
func NewGenerator(config Config) (*Generator, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, RSA_BITLEN)
	if err != nil {
		return nil, err
	}
	return &Generator{
		privateKey: privateKey,
		config:     config,
	}, nil
}

type Generator struct {
	privateKey *rsa.PrivateKey
	config     Config
}

// WritePrivateKeyPEMWithPW write a AES128 encrypt ASCII PEM key to filepath.
func (g *Generator) WritePrivateKeyPEMWithPW(filepath string, password string) error {
	pkder := x509.MarshalPKCS1PrivateKey(g.privateKey)
	block, err := x509.EncryptPEMBlock(rand.Reader,
		PEM_RSA_PRIVATE_KEY_TYPE, pkder,
		[]byte(password), x509.PEMCipherAES128)

	if err != nil {
		return err
	}

	d := pem.EncodeToMemory(block)
	return ioutil.WriteFile(filepath, d, 0600)
}

// WritePrivateKeyPEM write a plain ASCII PEM key to filepath.
func (g *Generator) WritePrivateKeyPEM(filepath string) error {
	pkder := x509.MarshalPKCS1PrivateKey(g.privateKey)
	d := encodePEM(PEM_RSA_PRIVATE_KEY_TYPE, pkder)
	return ioutil.WriteFile(filepath, d, 0600)
}

// WriteCertificatePEM write a ASCII PEM certificate to filepath.
func (g *Generator) WriteCertificatePEM(filepath string) error {
	now := time.Now()
	// TODO: (maybe )add other subject fields.
	subject := pkix.Name{
		CommonName: g.config.CommonName,
	}

	certTemp := &x509.Certificate{
		SerialNumber:          big.NewInt(now.UnixNano()),
		Subject:               subject,
		Issuer:                subject,
		BasicConstraintsValid: true,
		IsCA:      true,
		NotBefore: now,
		NotAfter:  now.Add(time.Duration(g.config.Days) * 24 * time.Hour),
	}

	if len(g.config.DNSNames) > 0 {
		certTemp.DNSNames = g.config.DNSNames
	}

	if len(g.config.IPAddresses) > 0 {
		certTemp.IPAddresses = g.config.IPAddresses
	}

	crtder, err := x509.CreateCertificate(rand.Reader, certTemp, certTemp, g.privateKey.Public(), g.privateKey)

	if err != nil {
		return err
	}

	d := encodePEM(PEM_CERTIFICATE_TYPE, crtder)
	return ioutil.WriteFile(filepath, d, 0600)
}

func encodePEM(PEMType string, derData []byte) []byte {
	block := &pem.Block{
		Type:    PEMType,
		Headers: nil,
		Bytes:   derData,
	}
	return pem.EncodeToMemory(block)
}
