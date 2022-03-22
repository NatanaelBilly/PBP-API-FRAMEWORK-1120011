package controllers

import (
	"log"
)

func GetUsersByName(name string) User {
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Users WHERE name=?",
		name,
	)
	if err != nil {
		log.Print(err)
	}

	var user User
	var users []User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Address, &user.Username, &user.Password); err != nil {
			log.Print(err.Error())
		} else {
			users = append(users, user)
		}
	}
	return (user)
}

func GetUsersByLogin(cekLogin Login) User {
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Users WHERE username=? AND password=?", cekLogin.Username, cekLogin.Password)
	if err != nil {
		log.Print(err)
	}

	var user User
	var users []User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Address, &user.Username, &user.Password); err != nil {
			log.Print(err.Error())
		} else {
			users = append(users, user)
		}
	}
	return (user)
}

func AddUser(cekRegister Register) bool {
	db := connect()
	defer db.Close()

	_, errQuery := db.Exec("INSERT INTO users(name, age, address, username, password) values (?,?,?,?,?)", cekRegister.Name, cekRegister.Age, cekRegister.Address, cekRegister.Username, cekRegister.Password1)

	if errQuery == nil {
		return true
	} else {
		return false
	}
}

func UpdateUser(user User) bool {
	db := connect()
	defer db.Close()

	_, errQuery := db.Exec("UPDATE users SET name = ?, username = ?, age = ?, address = ? WHERE id = ?", user.Name, user.Username, user.Age, user.Address, user.Id)
	if errQuery == nil {
		return true
	} else {
		return false
	}
}

func DeleteUser(id int) bool {
	db := connect()
	defer db.Close()

	_, errQuery := db.Exec("DELETE FROM users WHERE id = ?", id)

	if errQuery == nil {
		return true
	} else {
		return false
	}
}
