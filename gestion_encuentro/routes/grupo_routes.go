package routes

import(
	"gestion_encuentro/controllers"
	
	"github.com/gorilla/mux"
)

func RegisterGrupoRoutes(r *mux.Router){
	r.HandleFunc("/grupo", controllers.GetALLGrupos).Methods("GET")
	r.HandleFunc("/grupo/{id}", controllers.GetGrupoByID).Methods("GET")
	//r.HandleFunc("/grupo", controllers.CreateGrupo).Methods("POST")
	//r.HandleFunc("/grupo/{id}", controllers.UpdateGrupo).Methods("PUT")
	//r.HandleFunc("/grupo/{id}", controllers.DeleteGrupo).Methods("DELETE")
}