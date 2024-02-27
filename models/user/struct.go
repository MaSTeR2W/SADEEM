package user

import "github.com/MaSTeR2W/SADEEM/models/classifications"

type User struct {
	UserId          int                              `json:"user_id"`
	Name            string                           `json:"name"`
	Email           string                           `json:"email"`
	Image           string                           `json:"image"`
	UserType        string                           `json:"user_type"`
	Classifications []classifications.Classification `json:"classifications"`
}
