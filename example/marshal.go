package example

import "github.com/othon-hugo/jwt"

func Marshal() any {
	h := jwt.HeaderParams{
		Alg: jwt.HS256,
		Typ: "JWT",
	}

	claims := Info{ID: 1, Username: "foobar"}

	token, err := jwt.Marshal(h, claims, Secret)

	if err != nil {
		return err
	}

	return token
}
