package storage

import "errors"

type Profile struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type Game struct {
	Users []string `json:"users,omitempty"`
}

type Storage interface {
	Connect() error
	Close() error

	CreateGame(game Game) (int, error)

	CheckProfile(profile Profile) (bool, error)
	CheckGame(id int, user string) (bool, error)
	//CreateProfile(profile Profile) error
}

var (
	ErrProfileNotFound = errors.New("profile not found\n")
)
