package Model

import "time"

type RefreshToken struct {
	ID        int `json:"id"`
	UserID    int `json:"userID"`
	Token     int
	JwtID     int
	IsUsed    bool
	IsRevoked bool
	IssueAt   bool
	ExpireAt  time.Time
}
