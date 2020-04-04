package app

import (
	"net/http"

	"github.com/sarmad1995/mvc/controllers"
)

// StartApp is a function to start the app
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
