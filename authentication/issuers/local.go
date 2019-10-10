package issuers

import (
	"crypto/rsa"
	"errors"
	"io"
)

type Local struct {
	PublicKey  crypto.PublicKey
	PrivateKey crypto.PrivateKey
}

func (issuer *Local) Public() crypto.PublicKey {
	return issuer.Public
}

func (issuer *Local) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	rsaKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key failed rsa.PrivateKey type assertion")
	}

	return rsa.SignPKCS1v15(rand, rsaKey, opts.HashFunc(), digest)
}

func (issuer *Local) Verify(signature []byte, digest []byte, opts crypto.SignerOptions) error {
	rsaKey, ok := issuer.Public().(*rsa.PublicKey)
	if !ok {
		return errors.New("public key failed rsa.PublicKey type assertion")
	}

	return rsa.VerifyPKCS1v15(rsaKey, opts.HashFunc(), digest, signature)
}
