package main

//package insecurecert

import (
	"encoding/pem"
	"io/ioutil"
	"crypto/tls"
	"log"
	"os"
	"strconv"
)

type Cert struct {
	host string
	port int
}

func (c Cert) addr() string {
	return c.host + ":" + strconv.Itoa(c.port)
}

func (c Cert) der() ([]byte, error) {
	conf := tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", c.addr(), &conf)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	der := conn.ConnectionState().PeerCertificates[0].Raw

	return der, nil
}

func main() {

	c := Cert{
		host: "self-signed.badssl.com",
		port: 443,
	}

	d, derErr := c.der()
	if derErr != nil {
		log.Fatalf("Failed to download der: %s", derErr)
	}

	ioutil.WriteFile("key.der", d, 0644)


	certOut, err := os.Create("cert.pem")
	if err != nil {
		log.Fatalf("Failed to open cert.pem for writing: %s", err)
	}
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: d}); err != nil {
		log.Fatalf("Failed to write data to cert.pem: %s", err)
	}
}
