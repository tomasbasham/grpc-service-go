import (
	"crypto"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/tomasbasham/grpc-service-go/authentication/issuers"
)

const (
	encoder       = base64.RawURLEncoding
	numSeparators = 2
	jwtSeparator  = "."
)

type Issuer interface {
	Public() (crypto.PublicKey, error)
	Private() (crypto.PrivateKey, error)
}

func EncodeWithSigner(claim *ClaimSet, signer crypto.Signer, opts ...SignerOption) ([]byte, error) {
	h := header{Type: "JWT"}
	for _, opt := range opts {
		opt.Apply(&h)
	}

	hs, err := json.Marshal(h)
	if err != nil {
		return nil, fmt.Errorf("could not marshal header: %v", err)
	}

	cs, err := json.Marshal(claim)
	if err != nil {
		return nil, fmt.Errorf("could not marshal claim: %v", err)
	}

	// Length (in bytes) of each section in the JWT.
	headerLength := encoder.EncodedLen(len(hs))
	claimLength := encoder.EncodedLen(len(cs))
	sigLength := 1 // Not sure how to get this yet

	// Create a fixed length byte buffer to store the token.
	token := make([]byte, headerLength+claimsLength+sigLength+numSeparators)

	encoder.Encode(token, hs)
	token[headerLength] = jwtSeparator
	encoder.Encode(token[headerLength+1:], cs)

	h.Write(token[:headerLength+claimsLength])
	sig, err := signer.Sign(rand.Reader, hash.Sum(nil), hash)
	if err != nil {
		fmt.Errorf("unable to sign JWT: %v", err)
	}

	signedLength := headerLength + claimLength
	token[headerLength+1+claimsLength] = jwtSeparator
	jwt := fmt.Sprintf("%s.%s", ss, base64.RawURLEncoding.EncodeToString(sig))
	return []byte(jwt), nil
}

func Encode(claim *ClaimSet, key crypto.PrivateKey) ([]byte, error) {
	signer := HMACSigner{
		privateKey: key,
	}

	return EncodeWithSigner(claim, signer, WithHashFunction(crypto.SHA256))
}

// ClaimSet
type ClaimSet struct {
	Issuer     string `json:"iss"`             // email address of the client_id of the application making the access token request
	Scope      string `json:"scope,omitempty"` // space-delimited list of the permissions the application requests
	Audience   string `json:"aud"`             // descriptor of the intended target of the assertion (Optional).
	Expiration int64  `json:"exp"`             // the expiration time of the assertion (seconds since Unix epoch)
	IssuedAt   int64  `json:"iat"`             // the time the assertion was issued (seconds since Unix epoch)
	Typ        string `json:"typ,omitempty"`   // token type (Optional).

	// Email for which the application is requesting delegated access (Optional).
	Subject string `json:"sub,omitempty"`
}

// MarshalJSON
// func (cs *ClaimSet) MarshalJSON() ([]byte, error) {
// 	return nil, nil
// }

// VerifyWithSigner
func VerifyWithSigner(token []byte, signer Signer) error {
	parts := strings.Split(string(token), ".")
	if len(parts) != 3 {
		return errors.New("jws: invalid token received, token must have 3 parts")
	}

	h := sha256.New()
	h.Write([]byte(""))

	signer.Verify(sig, h.Sum(nil), crypto.SHA256)
}

// Verify
func Verify(publicKey crypto.PublicKey, token []byte) error {
	issuer := issuers.Local{
		PublicKey: publicKey,
	}

	return VerifyWithSigner(token, issuer)
}
