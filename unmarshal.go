package jwt

func Marshal(headerParams HeaderParams, claims any, secret []byte) (string, error) {
	return (&token{header: headerParams, payload: payload{claims: claims}}).marshal(secret)
}
