package encryption

import (
	"fmt"

	"github.com/billgraziano/dpapi"
)

func Decrypt(b []byte) []byte {
	decrypted, err := dpapi.DecryptBytes(b)
	if err != nil {
		fmt.Printf("Err in Decrypt: %s\n", err)
	}
	return decrypted
}

func Encrypt(b []byte) []byte {
	encrypted, err := dpapi.EncryptBytes(b)
	if err != nil {
		fmt.Printf("Err in Encrypt: %s\n", err)
	}
	return encrypted
}
