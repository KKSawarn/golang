package routers

import (
	"usercurd/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	userRoute := mux.NewRouter()

	userRoute.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	userRoute.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	userRoute.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	userRoute.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	userRoute.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	// userRoute.HandleFunc("/users", user.GetUsers).Methods("GET")
	// userRoute.HandleFunc("/user/{id}", user.GetUser).Methods("GET")
	// userRoute.HandleFunc("/user", user.CreateUser).Methods("POST")
	// userRoute.HandleFunc("/user/{id}", user.UpdateUser).Methods("PUT")
	// userRoute.HandleFunc("/user/{id}", user.DeleteUser).Methods("DELETE")

	return userRoute
}
