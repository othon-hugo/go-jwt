package jwt

func Marshal(h HeaderData, claims any, secret []byte) (string, error) {
	return (&token{header: h.header(), payload: payload{claims: claims}}).marshal(secret)
}
