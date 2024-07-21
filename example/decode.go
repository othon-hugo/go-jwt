package example

import "github.com/othon-hugo/jwt"

func Unmarshal() any {
	token, ok := Marshal().(string)

	if !ok {
		return ok
	}

	claims := Info{}

	err := jwt.Unmarshal(token, &claims, Secret)

	if err != nil {
		return err
	}

	return claims
}
