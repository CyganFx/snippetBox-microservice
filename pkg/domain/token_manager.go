package domain

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenManager interface {
	NewJWT(user *User) (string, error)
}

type Manager struct {
	signingKey string
}

func NewManager(signingKey string) *Manager {
	return &Manager{signingKey: signingKey}
}

func (m *Manager) NewJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      email,
		"authorized": true,
		"exp":        time.Now().Add(time.Hour).Unix(),
	})
	return token.SignedString([]byte(m.signingKey))
}
