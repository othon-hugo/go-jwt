package jwt

func Decode(encodedToken string, claims any, secret []byte) error {
	return (&token{payload: payload{claims: claims}}).unmarshal(encodedToken, secret)
}
