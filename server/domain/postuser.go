package domain

import "time"

type PostUser struct {
	UUID      string    `db:"uuid"`
	UserID    int64     `db:"user_id"`
	AccountID string    `db:"account_id"`
	Username  string    `db:"username"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
