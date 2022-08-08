package entities

type Category struct {
	ID   int64  `db:"categories.id" json:"id,omitempty"`
	Name string `db:"categories.name" json:"name,omitempty"`
}
