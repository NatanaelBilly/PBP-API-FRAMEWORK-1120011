package controllers

type User struct {
	Id       int    `form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	Age      int    `form:"age" json:"age"`
	Address  string `form:"address" json:"address"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type Login struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type Register struct {
	Name      string `form:"name" json:"name"`
	Age       int    `form:"age" json:"age"`
	Address   string `form:"address" json:"address"`
	Username  string `form:"username" json:"username"`
	Password1 string `form:"password1" json:"password1"`
	Password2 string `form:"password2" json:"password2"`
}

type Update struct {
	Name     string `form:"name" json:"name"`
	Age      int    `form:"age" json:"age"`
	Address  string `form:"address" json:"address"`
	Username string `form:"username" json:"username"`
}

type Delete struct {
	Id int `form:"id" json:"id"`
}
