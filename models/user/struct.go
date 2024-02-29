package user

import (
	"strconv"

	"github.com/MaSTeR2W/SADEEM/models/classification"
)

type User struct {
	UserId          int                            `json:"userId" db:"user_id"`
	Name            string                         `json:"name"`
	Email           string                         `json:"email"`
	Image           string                         `json:"image"`
	UserType        string                         `json:"userType" db:"user_type"`
	Classifications classification.Classifications `json:"classification"`
}

func (u *User) StrUserId() string {
	return strconv.FormatInt(int64(u.UserId), 10)
}
