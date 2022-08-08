package entities

import "time"

type BlogPost struct {
	ID          int64     `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Content     string    `db:"content" json:"content"`
	UserID      int64     `db:"userId" json:"userId"`
	Published   time.Time `db:"published" json:"published"`
	Updated     time.Time `db:"updated" json:"updated"`
	CategoryIDs []int64   `json:"categoryIds"`
}

type BlogPostResponse struct {
	ID         int64               `db:"id" json:"id"`
	Title      string              `db:"title" json:"title"`
	Content    string              `db:"content" json:"content"`
	UserID     int64               `db:"userId" json:"userId"`
	Published  time.Time           `db:"published" json:"published"`
	Updated    time.Time           `db:"updated" json:"updated"`
	User       UserWithoutPassword `db:"user" json:"user"`
	Categories []Category          `db:"categories" json:"categories"`
}
