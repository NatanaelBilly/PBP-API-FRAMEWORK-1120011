package main

import (
	controller "PBP-API-FRAMEWORK-1120011/controllers"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(http.StatusOK, "index", nil)
	})

	m.Post("/user", binding.Bind(controller.Login{}), func(r render.Render, cekLogin controller.Login) {
		var user controller.User
		user = controller.GetUsersByLogin(cekLogin)

		r.HTML(http.StatusOK, "user_info", user)
	})

	m.Post("/register", binding.Bind(controller.Register{}), func(r render.Render, cekRegister controller.Register) {
		var cek bool
		if cekRegister.Password1 == cekRegister.Password2 {
			cek = controller.AddUser(cekRegister)
			if cek == true {
				r.HTML(http.StatusOK, "registerSuccess", nil)
				r.HTML(http.StatusOK, "index", nil)
			} else {
				r.HTML(http.StatusOK, "registerFailed", nil)
				r.HTML(http.StatusOK, "index", nil)
			}
		} else {
			r.HTML(http.StatusOK, "registerFailed", nil)
			r.HTML(http.StatusOK, "index", nil)
		}
	})

	m.Post("/edit", binding.Bind(controller.Update{}), func(r render.Render, update controller.Update) {
		var user controller.User
		var temp controller.User
		user = controller.GetUsersByName(update.Name)
		temp = user
		temp.Name = update.Name
		temp.Username = update.Username
		temp.Age = update.Age
		temp.Address = update.Address

		var cek bool
		cek = controller.UpdateUser(user)
		if cek {
			r.HTML(http.StatusOK, "updatesuccess", nil)
			r.HTML(http.StatusOK, "user_info", temp)
		}
	})

	m.Post("/delete", binding.Bind(controller.Delete{}), func(r render.Render, delete controller.Delete) {
		var cek bool
		cek = controller.DeleteUser(delete.Id)
		if cek {
			r.HTML(http.StatusOK, "deletesuccess", nil)
			r.HTML(http.StatusOK, "index", nil)
		}
	})

	m.Get("/user/:username", func(r render.Render, p martini.Params) {
		var name string
		var user controller.User

		name = p["username"]
		user = controller.GetUsersByName(name)
		r.HTML(http.StatusOK, "user_info", user)
	})
}
