package example

type Info struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

var SecretKey = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
