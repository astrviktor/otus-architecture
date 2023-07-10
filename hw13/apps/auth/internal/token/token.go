package token

import "github.com/cristalhq/jwt/v5"

const secret = "secret"

type userClaims struct {
	GameID     int    `json:"gameId"`
	User       string `json:"user"`
	Operations []int  `json:"operations"`
}

func GetToken(id int, user string) (string, error) {
	key := []byte(`secret`)
	signer, err := jwt.NewSignerHS(jwt.HS256, key)

	if err != nil {
		return "", err
	}

	claims := &userClaims{
		GameID:     id,
		User:       user,
		Operations: []int{1, 2, 3},
	}

	// create a Builder
	builder := jwt.NewBuilder(signer)

	// and build a Token
	token, err := builder.Build(claims)
	if err != nil {
		return "", err
	}

	return token.String(), nil
}
