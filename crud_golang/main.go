package main

import (
	db "crud/DB"
	"crud/controller"
	"log"
	"net/http"
)

func main() {
	db.ConnectDB()
	// controller.Insert("laki-laki")
	// controller.Update(7, "wanita")
	// controller.Delete(7)
	controller.Select()

	log.Println("server runing on port 8080 ")
	http.ListenAndServe(":8080", nil)
}
