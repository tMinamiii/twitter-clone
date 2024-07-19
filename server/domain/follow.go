package domain

import "time"

type Follow struct {
	ID           int64     `db:"id"`
	UserID       int64     `db:"user_id"`
	FollowUserID int64     `db:"follow_user_id"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
