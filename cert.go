package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"strconv"
)

type Cert struct {
	host string
	port int
}

func (c Cert) addr() string {
	return c.host + ":" + strconv.Itoa(c.port)
}

func (c Cert) Der() ([]byte, error) {
	conf := tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", c.addr(), &conf)
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	return conn.ConnectionState().PeerCertificates[0].Raw, nil
}

func main() {

	host := "untrusted-root.badssl.com"
	//host := "self-signed.badssl.com"

	c := Cert{
		//host: "untrusted-root.badssl.com",
		host: host,
		port: 443,
	}

	certDer, derErr := c.Der()
	if derErr != nil {
		log.Fatalf("Failed to download der: %s", derErr)
	}

	ioutil.WriteFile(host + "-byte2.der", certDer, 0644)
}
