package main

//package insecurecert

import (
	"io/ioutil"
	"crypto/tls"
	"strconv"
)

type Cert struct {
	host string
	port int
}

func (c Cert) addr() string {
	return c.host + strconv.Itoa(c.port)
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

	d, _ := c.der()

	ioutil.WriteFile("key.der", d, 0644)
}
