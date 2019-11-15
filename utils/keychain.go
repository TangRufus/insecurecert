package utils

import (
	"log"
)

type Keychain struct{}

func (Keychain) Trust(certificatePath string, resultType string) error {
	const keychainPath = "/Library/Keychains/System.keychain"
	log.Printf("Adding (%s) %s to the admin cert store in system keychain %s", resultType, certificatePath, keychainPath)

	args := []string{
		"add-trusted-cert",
		"-d",
		"-r", resultType,
		"-k", keychainPath,
		certificatePath,
	}

	return Cmd{}.Run("security", args...)
}
