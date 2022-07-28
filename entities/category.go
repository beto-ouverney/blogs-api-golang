package entities

type Category struct {
	ID   int64  `db:"id" json:"id,omitempty"`
	Name string `db:"name" json:"name,omitempty"`
}
