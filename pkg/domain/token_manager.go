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

func (m *Manager) NewJWT(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID,
		"username":   user.Name,
		"authorized": true,
		"exp":        time.Now().Add(time.Hour).Unix(),
	})
	return token.SignedString([]byte(m.signingKey))
}
