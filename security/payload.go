package security

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrorInvalidToken = errors.New("Token invalido")
	ErrorExpiredToken = errors.New("Token expirado")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Role      string    `json:"role"`
	Name      string    `json:"name"`
	Image     string    `json:"image_url"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(role string, name string, image string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		Name:      name,
		Role:      role,
		Image:     image,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrorExpiredToken
	}
	return nil
}
