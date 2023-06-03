package token

import (
	"encoding/json"
	"github.com/cristalhq/jwt/v5"
	"server/internal/handlers"
)

const secret = "secret"

type userClaims struct {
	GameID     int    `json:"gameId"`
	User       string `json:"user"`
	Operations []int  `json:"operations"`
}

func CheckToken(message handlers.Message) (bool, error) {
	key := []byte(`secret`)
	verifier, err := jwt.NewVerifierHS(jwt.HS256, key)

	if err != nil {
		return false, err
	}

	newToken, err := jwt.Parse([]byte(message.Token), verifier)
	if err != nil {
		return false, err
	}

	var newClaims userClaims
	err = json.Unmarshal(newToken.Claims(), &newClaims)
	if err != nil {
		return false, err
	}

	if message.GameID != newClaims.GameID {
		return false, nil
	}

	if message.User != newClaims.User {
		return false, nil
	}

	for _, o := range newClaims.Operations {
		if o == message.OperationID {
			return true, nil
		}
	}

	return false, nil
}
