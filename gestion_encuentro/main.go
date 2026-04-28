package main

import (
	"log"
	"net/http"

	"API_GO_CRUD/config"
	"API_GO_CRUD/routes"

	"github.com/gorilla/mux"
)

// middleware CORS
func enableCORS(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		// Permiten cualquier origen de la petición
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		if r.Method == "OPTIONS" {
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

func main(){
	config.ConnectDB() // Conexión a la base de datos

	r := mux.NewRouter()

	// Registro de rutas
	//routes.
	//routes.
	//routes.
	//routes.
	//routes.

	log.Println("Servidor operando en el puerto 8090")

	log.Fatal(http.ListenAndServe(":8090", enableCORS(r)))
}