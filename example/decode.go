package example

import "github.com/othon-hugo/jwt"

func Decode() any {
	token, ok := Encode().(string)

	if !ok {
		return ok
	}

	claims := Info{}

	err := jwt.Decode(token, &claims, Secret)

	if err != nil {
		return err
	}

	return claims
}
