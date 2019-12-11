package utils

import (
	"log"
)

type Browser struct {}

func (Browser) Open (url string) error {
	log.Printf("Opening %s in browser...", url)
	return Cmd{}.Run("open", url)
}
