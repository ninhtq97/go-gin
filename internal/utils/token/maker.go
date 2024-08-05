package token

import (
	"ninhtq/go-gin/core/exception"
	"time"

	"github.com/google/uuid"
)

type Maker interface {
	CreateToken(userID string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(userID string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, exception.New(exception.TypeInternal, "failed generate token id", err)
	}
	payload := &Payload{
		ID:        tokenID,
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return exception.New(exception.TypeTokenExpired, "token expired", nil)
	}
	return nil
}
