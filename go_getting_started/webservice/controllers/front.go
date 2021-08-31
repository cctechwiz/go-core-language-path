package controllers

import "net/http"

func RegisterControllers() {
	us := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
