package main

import (
	"gestion_encuentro/routes"
	"gestion_encuentro/config"
	"log"
	"net/http"

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
	routes.RegisterEncuentroRoutes(r)
	routes.RegisterGrupoEncuentroRoutes(r)
	//routes.RegisterGrupoEquipoRoutes(r)
	routes.RegisterGrupoRoutes(r)

	log.Println("Servidor operando en el puerto 8090")

	log.Fatal(http.ListenAndServe(":8090", enableCORS(r)))
}