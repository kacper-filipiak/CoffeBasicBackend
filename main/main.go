package main

import (
	"fmt"
	"kacperfilipiak/coffe_home/api"
	"kacperfilipiak/coffe_home/db"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("server is up")
	db_connection, err := db.StartDB()
	if err != nil {
		log.Println("error while connecting to db: " + err.Error())
	} else {
		log.Println("Created db" + db_connection.String())
	}
	router := api.StartApi(db_connection)
	port := os.Getenv("PORT")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}

}
