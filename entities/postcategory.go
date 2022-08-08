package entities

type PostCategory struct {
	PostID     int64 `db:"postId" json:"postId,omitempty"`
	CategoryID int64 `db:"categoryId" json:"categoryId,omitempty"`
}
