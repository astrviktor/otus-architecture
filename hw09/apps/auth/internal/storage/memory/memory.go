package memory

import (
	"auth/internal/storage"
	"sync"
)

type Storage struct {
	profiles map[string]storage.Profile
	lastGame int
	games    map[int]storage.Game
	mutex    *sync.Mutex
}

func New() *Storage {
	mutex := sync.Mutex{}

	s := &Storage{
		profiles: make(map[string]storage.Profile),
		lastGame: 0,
		games:    make(map[int]storage.Game),
		mutex:    &mutex,
	}

	s.Init()

	return s
}

func (s *Storage) Init() {
	s.mutex.Lock()

	s.profiles["user1"] = storage.Profile{
		Username: "user1",
		Password: "pass1",
	}

	s.profiles["user2"] = storage.Profile{
		Username: "user2",
		Password: "pass2",
	}

	s.profiles["user3"] = storage.Profile{
		Username: "user3",
		Password: "pass3",
	}

	s.mutex.Unlock()
}

func (s *Storage) Connect() error {
	return nil
}

func (s *Storage) Close() error {
	return nil
}

func (s *Storage) CreateGame(game storage.Game) (int, error) {
	s.mutex.Lock()
	s.lastGame++
	n := s.lastGame
	s.games[n] = game
	s.mutex.Unlock()

	return n, nil
}

func (s *Storage) CheckProfile(profile storage.Profile) (bool, error) {
	s.mutex.Lock()
	p, ok := s.profiles[profile.Username]
	s.mutex.Unlock()

	if !ok {
		return false, nil
	}

	if p.Password != profile.Password {
		return false, nil
	}

	return true, nil
}

func (s *Storage) CheckGame(id int, user string) (bool, error) {
	s.mutex.Lock()
	game, ok := s.games[id]
	s.mutex.Unlock()

	if !ok {
		return false, nil
	}

	for _, u := range game.Users {
		if u == user {
			return true, nil
		}
	}

	return false, nil
}
