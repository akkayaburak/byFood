package models

type User struct {
	UserID       string `json:"userid"`
	UserName     string `json:"username"`
	Password     string `json:"password"`
	PasswordHash string `json:"passwordhash"`
	Mail         string `json:"mail"`
	IsDeleted    bool   `json:"isdeleted"`
}

type UserID struct {
	UserID string `json:"userid"`
}
