package example

import "github.com/othon-hugo/jwt"

func Unmarshal() any {
	token, ok := Marshal().(string)

	if !ok {
		return ok
	}

	claims := Info{}

	if err := jwt.Unmarshal(token, &claims, SecretKey); err != nil {
		return err
	}

	return claims
}
