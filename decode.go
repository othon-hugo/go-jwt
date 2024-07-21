package jwt

func Unmarshal(encodedToken string, claims any, secret []byte) error {
	return (&token{payload: payload{claims: claims}}).unmarshal(encodedToken, secret)
}
