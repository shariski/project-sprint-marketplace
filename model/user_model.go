package model

type UserModel struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserGetModel struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}
