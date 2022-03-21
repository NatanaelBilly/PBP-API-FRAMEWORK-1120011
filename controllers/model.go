package controllers

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
	Username  string `json:"username"`
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
}

type Update struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Username string `json:"username"`
}

type Delete struct {
	Id int `json:"id"`
}
