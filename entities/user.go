import "time"

//User presents a user entity.
type User struct {
	ID          int        `db: "id" json:"id"`
	DisplayName string     `db: "display_name" json:"displayName"`
	Email       string     `db: "email" json:"email"`
	Password    string     `db: "password" json:"password"`
	Image       string     `db: "image" json:"image"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}
