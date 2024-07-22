package test

import (
	"testing"

	"github.com/othon-hugo/jwt"
)

func TestMarshalExpectedTokensMatch(t *testing.T) {
	for i := 0; i < len(Claims); i++ {
		claims, expectedEncodedToken := Claims[i], HS256Tokens[i]

		h := jwt.HeaderParams{
			Alg: jwt.HS256,
			Typ: "JWT",
		}

		encodedToken, err := jwt.Marshal(h, claims, SecretKey)

		if err != nil {
			t.Fatalf("JWT encoding error: %v", err)
		}

		if encodedToken != expectedEncodedToken {
			t.Errorf("Token comparison failed: Got %s, Expected: %s", encodedToken, expectedEncodedToken)
		}
	}
}

func TestUnmarshalExpectedTokensMatch(t *testing.T) {
	for i := 0; i < len(Claims); i++ {
		expectedClaims, encodedToken := Claims[i], HS256Tokens[i]

		resultedClaims := UserInfo{}

		if err := jwt.Unmarshal(encodedToken, &resultedClaims, SecretKey); err != nil {
			t.Fatalf("JWT encoding error: %v", err)
		}

		if resultedClaims.ID != expectedClaims.ID || resultedClaims.Name != expectedClaims.Name {
			t.Errorf("Token comparison failed: Got %v, Expected: %v", resultedClaims, expectedClaims)
		}
	}
}

func TestUnmarshalInvalidSecretReturnError(t *testing.T) {
	invalidSecret := []byte("invalid-secret-key")

	for i := 0; i < len(Claims); i++ {
		token := HS256Tokens[i]

		decodedPayload := UserInfo{}

		if err := jwt.Unmarshal(token, &decodedPayload, invalidSecret); err == nil {
			t.Error("Expected error, but none occurred")
		}
	}
}
