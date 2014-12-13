package main

import (
	"darling"
	"log"
	"plant/controllers"
)

func main() {
	app := darling.NewApp()
	app.Handlers.Add("/users/([0-9]+)", &controllers.UserController{})
	app.Handlers.Add("/join", &controllers.RegisterController{})
	app.Server.Addr = ":8001"
	app.Server.Handler = app.Handlers
	err := app.Server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}
