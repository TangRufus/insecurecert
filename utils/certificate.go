package utils

import (
	"crypto/tls"
	"strconv"
)

type Certificate struct {
	Hostname string
	Port     int
}

func (c Certificate) Addr() string {
	return c.Hostname + ":" + strconv.Itoa(c.Port)
}

func (c Certificate) Der() ([]byte, error) {
	conf := tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", c.Addr(), &conf)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return conn.ConnectionState().PeerCertificates[0].Raw, nil
}
