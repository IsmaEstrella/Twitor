package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/IsmaEstrella/Twitor/middlew"
	"github.com/IsmaEstrella/Twitor/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.RevisionBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.RevisionBD(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "2300"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) //Se pone a escuchar al servidor, y handler le va asigna el puerto 8080
}
