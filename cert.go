package main

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"crypto/tls"
	"log"
	//"os"
	"strconv"
)

type Cert struct {
	host string
	port int
}

func (c Cert) addr() string {
	return c.host + ":" + strconv.Itoa(c.port)
}

func (c Cert) der() ([]*x509.Certificate, error) {
	conf := tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", c.addr(), &conf)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return conn.ConnectionState().PeerCertificates, nil
}

func main() {

	c := Cert{
		//host: "untrusted-root.badssl.com",
		host: "self-signed.badssl.com",
		port: 443,
	}

	certs, derErr := c.der()
	if derErr != nil {
		log.Fatalf("Failed to download der: %s", derErr)
	}

	var certBytes []byte
	for _, cert := range certs {
		certBytes = append(certBytes, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw})...)
	}

	ioutil.WriteFile("self-signed.badssl.com-byte.der", certBytes, 0644)


	//certOut, err := os.Create("untrusted-root.pem")
	//if err != nil {
	//	log.Fatalf("Failed to open cert.pem for writing: %s", err)
	//}

}
