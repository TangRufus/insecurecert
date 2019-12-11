package command

import (
	"github.com/urfave/cli/v2"
	"insecurecert/utils"
	"io/ioutil"
	"log"
	"os"
)

type Trust struct{}

func (t Trust) Command() *cli.Command {
	return &cli.Command{
		Name:  "trust",
		Usage: "download certificate from website and add it to OS trust store",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "hostname",
				Usage:    "`HOSTNAME` to the website, without protocols (https://, http://, etc)",
				Required: true,
			},
			&cli.IntFlag{
				Name:  "port",
				Usage: "`PORT` to the website",
				Value: 443,
			},
			&cli.StringFlag{
				Name:  "result-type",
				Usage: "`RESULT_TYPE` must be one of trustRoot or trustAsRoot",
				Value: "trustRoot",
			},
		},
		Action: t.action(),
	}
}

func (Trust) action() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		hostname := c.String("hostname")
		port := c.Int("port")
		resultType := c.String("result-type")

		certificate := utils.Certificate{
			Hostname: hostname,
			Port:     port,
		}

		log.Printf("Downloading certificate from %s", certificate.Addr())

		derByte, derErr := certificate.Der()
		if derErr != nil {
			log.Println("Failed to download certificate")
			return derErr
		}

		derFile, tempFileErr := ioutil.TempFile("", "insecurecert.*.der")
		if tempFileErr != nil {
			log.Println("Failed to create temporary file")
			return tempFileErr
		}
		defer os.Remove(derFile.Name())

		// TODO: Refactor.
		// See: https://golang.org/pkg/io/ioutil/#WriteFile
		if _, err := derFile.Write(derByte); err != nil {
			derFile.Close()
			log.Println("Failed to write to temporary file")
			return err
		}
		// TODO: Refactor.
		// See: https://golang.org/pkg/io/ioutil/#WriteFile
		if err := derFile.Close(); err != nil {
			log.Println("Failed to close temporary file")
			return err
		}

		return (utils.Keychain{}).Trust(derFile.Name(), resultType)
	}
}
