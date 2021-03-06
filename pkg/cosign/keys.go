package cosign

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"

	"github.com/theupdateframework/go-tuf/encrypted"
)

type PassFunc func(bool) ([]byte, error)

type Keys struct {
	PrivateBytes []byte
	PublicBytes  []byte
}

func GenerateKeyPair(pf PassFunc) (*Keys, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	// Encrypt the private key and store it.
	password, err := pf(true)
	if err != nil {
		return nil, err
	}

	encBytes, err := encrypted.Encrypt(priv, password)
	if err != nil {
		return nil, err
	}

	privBytes := pem.EncodeToMemory(&pem.Block{
		Bytes: encBytes,
		Type:  "ENCRYPTED COSIGN PRIVATE KEY",
	})

	b, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, err
	}

	// Now do the public key
	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: b,
	})

	return &Keys{
		PrivateBytes: privBytes,
		PublicBytes:  pubBytes,
	}, nil
}
