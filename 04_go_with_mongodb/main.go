package main

import (
	"log"
	"net/http"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.Get("/", uc.GetAll)
	r.Post("/", uc.Create)
	r.Get("/:id", uc.GetUser)
	r.Put("/:id", uc.Update)
	r.Delete("/:id", uc.Delete)
	log.Fatal(http.ListenAndServe(":8080", r))

}
