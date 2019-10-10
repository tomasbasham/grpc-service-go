package authentication_test

import (
	"crypto"
	"crypto/rsa"
	"io"
	"testing"
)

func TestSignWithSigner(t *testing.T) {

}

func TestSign(t *testing.T) {

}

func TestVerifyWithSigner(t *testing.T) {

}

func TestVerify(t *testing.T) {

}

type TestSigner struct{}

func (issuer *TestSigner) Public() rsa.PublicKey {

}

func (issuer *TestSigner) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {

}

func (issuer *TestSigner) Verify(signature []byte, digest []byte, opts crypto.SignerOpts) error {

}
