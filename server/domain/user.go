package domain

import "time"

type User struct {
	ID        int64     `db:"id"`
	AccountID string    `db:"account_id"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
