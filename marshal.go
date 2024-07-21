package jwt

func Unmarshal(encodedToken string, claims any, secret []byte) error {
	tkn := &token{
		payload: payload{claims: claims},
	}

	if err := tkn.unmarshal(encodedToken, secret); err != nil {
		return err
	}

	if tkn.header.Typ != JWT {
		return UnsupportedTypeError{tkn.header.Typ}
	}

	return nil
}
