package main

import (
	"log"
	"net/http"
	"time"
	"usercurd/models"
	"usercurd/routers"
)

func main() {
	time.Sleep(15 * time.Second)
	models.InitialMigration()
	r := routers.SetupRouter()
	log.Println("server started on 9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}
