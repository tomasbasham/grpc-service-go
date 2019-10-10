package issuers

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"golang.org/x/oauth2/google"
)

// Google is a JWT issuer. This type implements `crypto.Signer` for it to be
// used to sign a JWT claim set.
type Google struct {
	serviceAccountKeyFile io.Reader
}

func (*Google) Public() (crypto.PublicKey, error) {
	return nil, nil
}

func (issuer *Google) Private() (crypto.PrivateKey, error) {
	sa, err := ioutil.ReadAll(issuer.serviceAccountKeyFile)
	if err != nil {
		return nil, fmt.Errorf("could not read service account file: %v", err)
	}

	conf, err := google.JWTConfigFromJSON(sa)
	if err != nil {
		return nil, fmt.Errorf("could not parse service account JSON: %v", err)
	}

	block, _ := pem.Decode(conf.PrivateKey)
	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("private key parse error: %v", err)
	}

	rsaKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key failed rsa.PrivateKey type assertion")
	}

	return rsaKey, nil
}

type RSASigner struct {
	publicKey  rsa.PublicKey
	privateKey rsa.PrivateKey
}

func NewRSASigner(issuer Issuer) RSASigner {
	publicKey, _ := issuer.Public()
	privateKey, _ := issuer.Private()

	return RSASigner{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func (s *RASSigner) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	return rsa.SignPKCS1v15(rand, s.privateKey, opts.HashFunc(), digest)
}

func (s *RSQSigner) Verify(signature []byte, digest []byte, opts crypto.SignerOptions) error {
	return rsa.VerifyPKCS1v15(s.publicKey, opts.HashFunc(), digest, signature)
}
