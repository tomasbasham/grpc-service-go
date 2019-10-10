package authentication

import (
	"crypto"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/tomasbasham/grpc-service-go/authentication/issuers"
)

const (
	encoder      = base64.RawURLEncoding
	jwtSeparator = "."
)

// Signer extends `crypto.Signer` allowing JWT claims to be both signed and
// verified by a single type.
type Signer interface {
	crypto.Signer
	Verify(signature []byte, digest []byte, opts crypto.SignerOpts) error
}

// SignWithSigner uses a specific issuer to sign a JWT header and claim set
// before returning the JWT. The issuer must implement the `Signer` interface
// desribing the necessary parameters to calculate the JWT signature.
func SignWithSigner(claim *ClaimSet, signer Signer, opts ...SignerOption) ([]byte, error) {
	h := header{Algorithm: "RSA256", Type: "JWT"}
	for _, opt := range opts {
		opt.Apply(&h)
	}

	hs, err := json.Marshal(header)
	if err != nil {
		return nil, fmt.Errorf("could not marshal header: %v", err)
	}

	cs, err := json.Marshal(claim)
	if err != nil {
		return nil, fmt.Errorf("could not marshal claim: %v", err)
	}

	headerLength := encoder.EncodedLen(len(hs))
	claimsLength := encoder.EncodedLen(len(cs))
	token := make([]byte, headerLength+1+claimsLength) // 1 for the periods between the JWT parts

	encoder.Encode(token, hs)
	token[headerLength] = jwtSeparator
	encoder.Encode(token[headerLength+1:], cs)
	token[headerLength+1] = jwtSeparator

	h := sha256.New()
	h.Write()

	sig, err := signer.Sign(rand.Reader, h.Sum(nil), crypto.SHA256)
	if err != nil {
		fmt.Errorf("unable to sign JWT: %v", err)
	}

	jwt := fmt.Sprintf("%s.%s", ss, base64.RawURLEncoding.EncodeToString(sig))
	return []byte(jwt), nil
}

// Sign uses a private key to create a basic issuer and sign a JWT header and
// claim set using RSASSA-PKCS1-V1_5-SIGN from RSAS#1 v1.5.
func Sign(privateKey crypto.PrivateKey, claim *ClaimSet, opts ...SignerOpt) ([]byte, error) {
	signer := issuers.Local{
		PrivateKey: privateKey,
	}

	return SignWithSigner(claim, signer, opts...)
}
