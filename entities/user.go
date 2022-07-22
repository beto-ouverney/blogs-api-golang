package entities

//User presents a user entity.
type User struct {
	ID          int    `db: "id" json:"id,omitempty"`
	DisplayName string `db: "displayName" json:"displayName,omitempty"`
	Email       string `db: "email" json:"email,omitempty"`
	Password    string `db: "password" json:"password,omitempty"`
	Image       string `db: "image" json:"image,omitempty"`
}

//User presents a user entity without password.
type UserWithoutPassword struct {
	ID          int    `db: "id" json:"id"`
	DisplayName string `db: "displayName" json:"displayName"`
	Email       string `db: "email" json:"email"`
	Password    string `db: "password" json:"password"`
	Image       string `db: "image" json:"image"`
}
