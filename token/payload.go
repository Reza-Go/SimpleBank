package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// payload data of token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAT  time.Time `json:"issued_at"`
	ExpiredAT time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAT:  time.Now(),
		ExpiredAT: time.Now().Add(duration),
	}

	return payload, nil
}

// check the payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAT) {
		return ErrExpiredToken
	}
	return nil
}
