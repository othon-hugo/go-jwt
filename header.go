package jwt

import (
	"hash"
	"strings"

	"github.com/othon-hugo/jwt/crypto"
	"github.com/othon-hugo/jwt/encoding"
)

const (
	HS256 = "HS256"
	HS384 = "HS384"
	HS512 = "HS512"
)

type HeaderParams = header

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

func (h *header) marshal() (string, error) {
	jsonHeader, err := encoding.EncodeJSON(h)

	if err != nil {
		return "", err
	}

	return encoding.EncodeJWTBase64(jsonHeader), nil
}

func (h *header) unmarshal(encodedHeader string) error {
	jsonHeader, err := encoding.DecodeJWTBase64(encodedHeader)

	if err != nil {
		return err
	}

	return encoding.DecodeJSON(jsonHeader, h)
}

func (h *header) signer(secret []byte) (hash.Hash, error) {
	switch strings.ToUpper(h.Alg) {
	case HS256:
		return crypto.NewHMAC(crypto.NewSHA256, secret), nil
	case HS384:
		return crypto.NewHMAC(crypto.NewSHA384, secret), nil
	case HS512:
		return crypto.NewHMAC(crypto.NewSHA512, secret), nil
	}

	return nil, UnsupportedAlgorithmError{h.Alg}
}
