package security

type AuthData struct {
	UserId   int    `db:"user_id"`
	UserType string `db:"user_type"`
}
